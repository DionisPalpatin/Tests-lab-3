from sql_faker.database import Database

def add_users(db, n = 100):

    db.add_table(table_name="users", n_rows=n)

    # db.tables["users"].add_primary_key(column_name="id")
    db.tables["users"].add_column(column_name="fio", data_type="text", data_target="name")
    db.tables["users"].add_column(column_name="registration_date", data_type="date_time", data_target="date_time")
    db.tables["users"].add_column(column_name="login", data_type="text", data_target="email")
    db.tables["users"].add_column(column_name="password", data_type="text", data_target="password")
    db.tables["users"].add_column(column_name="role", data_type="int", column_value=0)

    db.tables["users"].generate_data(recursive=False, lang="en_US")



db = Database(db_name="test_lab3")
add_users(db, 10000*5)


print(db.tables["users"].export_dml("ins.sql"))
