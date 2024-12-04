create schema if not exists test_lab3;


drop table if exists test.users;


create table test.users (
    id                serial       primary key,
    fio               varchar(255) not null,
    registration_date timestamptz  not null,
    login             varchar(255) not null unique,
    password          varchar(255) not null unique,
    role              int          not null check (role >= 0)
);


------------------------------------------------------------------------------------------------------------------------
-- Триггер на автовыдачу прав при регистрастрации или изменении данных пользователя
------------------------------------------------------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION main.update_user_role()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.role = 0 THEN
        -- Роль читателя (reader)
        EXECUTE format('ALTER ROLE user_%s SET ROLE Reader', NEW.id);
    ELSIF NEW.role = 1 THEN
        -- Роль автора (author)
        EXECUTE format('ALTER ROLE user_%s SET ROLE Author', NEW.id);
    ELSIF NEW.role = 2 THEN
        -- Роль администратора (admin)
        EXECUTE format('ALTER ROLE user_%s SET ROLE Admin', NEW.id);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;



DROP TRIGGER IF EXISTS update_user_role_trigger ON main.users;
CREATE TRIGGER update_user_role_trigger
    AFTER INSERT OR UPDATE ON main.users
    FOR EACH ROW
    EXECUTE PROCEDURE main.update_user_role();


------------------------------------------------------------------------------------------------------------------------
-- Trigger on table users
------------------------------------------------------------------------------------------------------------------------
create or replace function main.func_stat_user_trigger()
    returns trigger as $$
declare
    latest_stat_date timestamptz;
    current_time timestamptz := now();
begin
    -- Получаем максимальную дату из таблицы stat
    select max(stat_date) into latest_stat_date from main.stat;

    -- Копируем последнюю строку из таблицы main.stat
    insert into main.stat (total_users, total_readers, total_authors, total_admins, total_collections, total_teams, total_sections, total_notes, total_open_notes, total_close_notes, total_notes_in_collections, stat_date)
    select total_users, total_readers, total_authors, total_admins, total_collections, total_teams, total_sections, total_notes, total_open_notes, total_close_notes, total_notes_in_collections, current_time
    from main.stat
    order by stat_date desc
    limit 1;

    -- В зависимости от операции обновляем соответствующее поле
    if TG_OP = 'INSERT' then
        update main.stat
        set total_users = total_users + 1,
            total_readers = total_readers + case when new.role = 0 then 1 else 0 end,
            total_authors = total_authors + case when new.role = 1 then 1 else 0 end,
            total_admins = total_admins + case when new.role = 2 then 1 else 0 end
        where stat_date = latest_stat_date;

        return new;

    elsif TG_OP = 'DELETE' then
        update main.stat
        set total_users = total_users - 1,
            total_readers = total_readers - case when old.role = 0 then 1 else 0 end,
            total_authors = total_authors - case when old.role = 1 then 1 else 0 end,
            total_admins = total_admins - case when old.role = 2 then 1 else 0 end
        where stat_date = latest_stat_date;

        return old;

    elsif TG_OP = 'UPDATE' then
        -- Если изменяется роль пользователя
        if new.role != old.role then
            update main.stat
            set total_readers = total_readers + case when new.role = 0 then 1 else 0 end - case when old.role = 0 then 1 else 0 end,
                total_authors = total_authors + case when new.role = 1 then 1 else 0 end - case when old.role = 1 then 1 else 0 end,
                total_admins = total_admins + case when new.role = 2 then 1 else 0 end - case when old.role = 2 then 1 else 0 end
            where stat_date = latest_stat_date;

            return new;
        end if;
    end if;

    return new;
end;
$$ language plpgsql;

drop trigger if exists stat_user_trigger on main.users;
create trigger stat_user_trigger
    after insert or delete or update on main.users
    for each row
execute function main.func_stat_user_trigger();


------------------------------------------------------------------------------------------------------------------------
-- хранимая процедура для удаления пользователя и всех связанных данных
------------------------------------------------------------------------------------------------------------------------
create or replace procedure main.delete_user(deleted_user_id integer)
as $$
begin
	delete from main.notes_collections where collection_id in (select id from main.collections where owner_id = deleted_user_id);
	delete from main.collections where owner_id = deleted_user_id;
	delete from main.team_members where user_id = deleted_user_id;
	delete from main.texts where note_id in (select id from main.notes where owner_id = deleted_user_id);
	delete from main.images where note_id in (select id from main.notes where owner_id = deleted_user_id);
	delete from main.raw_datas where note_id in (select id from main.notes where owner_id = deleted_user_id);
	delete from main.notes where owner_id = deleted_user_id;
	delete from main.users where id = deleted_user_id;
end;
$$ language plpgsql;


------------------------------------------------------------------------------------------------------------------------
-- Роли
------------------------------------------------------------------------------------------------------------------------
drop role if exists Reader;
drop role if exists Author;
drop role if exists Administrator;

create role Reader login;
create role Author login;
create role Administrator login;

grant select on main.notes, main.texts, main.images, main.raw_datas, main.teams to Reader;
grant select on main.teams, main.team_members, main.sections, main.teams_sections to Reader;
grant select, update, insert, delete on main.collections, main.note_collections to Reader;
grant usage, select on all sequences in schema main to Reader;

grant select, update, insert, delete on main.notes, main.texts, main.images, main.raw_datas to Author;
grant select on main.teams, main.team_members, main.sections, main.teams_sections to Author;
grant select, update, insert, delete on main.collections, main.note_collections to Author;
grant usage, select on all sequences in schema main to Author;

grant select, update, insert, delete on main.notes, main.texts, main.images, main.raw_datas to Administrator;
grant select, update, insert, delete on main.collections, main.note_collections to Administrator;
grant select, update, insert, delete on main.teams_sections to Administrator;
grant select, update, insert, delete on main.teams to Administrator;
grant select, update, insert, delete on main.team_members to Administrator;
grant select, update, insert, delete on main.sections to Administrator;
grant select, update, insert, delete on main.users to Administrator;
grant usage, select on all sequences in schema main to Administrator;


------------------------------------------------------------------------------------------------------------------------
-- Суперадмин обязательно должен быть
------------------------------------------------------------------------------------------------------------------------
insert into main.users (fio, login, password, role) values
    ('mainadmin', 'mainadminlogin', 'mainadminpassword', 2);


------------------------------------------------------------------------------------------------------------------------
-- Прочие пользователи
------------------------------------------------------------------------------------------------------------------------