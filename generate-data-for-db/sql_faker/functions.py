from faker import Faker
import datetime
import pytz
import random


error_02 = "List contains to elements."


def get_fake_data(data_target="name", n_rows=100, lang="de_DE", timezone="Europe/Moscow"):
    """This method generates a certain data item based on the data_target selected
    
    :param data_target: Data item type that should be created
    :type data_target: String
    :param n_rows: Number of rows to generate
    :type n_rows: Integer
    :returns: List of data
    """
    data_faker = Faker(lang)
    generator_function = getattr(data_faker, data_target)
    return_list = []
    tz = pytz.timezone(timezone)
    domains = ['newmail.com', 'live.com', 'epost.com', 'myemail.com', 'coolmail.com', 'customdomain.com', 'icloud.com',
               'smtp.com', 'msn.com', 'fakemail.com', 'yourmail.com', 'message.com', 'mail.ru', 'digitalmail.com', 
               'custommail.com', 'zoho.com', 'yandex.com', 'directmail.com', 'gmx.com', 'inbox.com', 'yahoo.com', 'mailservice.com', 
               'message.net', 'example.com', 'ezyemail.com', 'randommail.com', 'mxmail.com', 'gmail.com', 'mailbox.com', 
               'mailinator.com', 'mail.com', 'hushmail.com', 'testmail.com', 'aol.com', 'service.com', 'supermail.com', 'online.com', 
               'webmail.com', 'protonmail.com', 'example.org', 'trashmail.com', 'tutanota.com', 'quickmail.com', 'hotmail.com', 
               'fastmail.com', 'outlook.com', 'inbox.net', 'mailhub.com', 'example.net', 'mymail.com', 'newmail.ru', 'live.ru', 'epost.ru', 
               'myemail.ru', 'coolmail.ru', 'customdomain.ru', 'icloud.ru', 'smtp.ru', 'msn.ru', 'fakemail.ru', 'yourmail.ru', 
               'message.ru', 'mail.ru', 'digitalmail.ru', 'custommail.ru', 'zoho.ru', 'yandex.ru', 'directmail.ru', 'gmx.ru', 'inbox.ru', 
               'yahoo.ru', 'mailservice.ru', 'message.net', 'example.ru', 'ezyemail.ru', 'randommail.ru', 'mxmail.ru', 'gmail.ru', 
               'mailbox.ru', 'mailinator.ru', 'mail.ru', 'hushmail.ru', 'testmail.ru', 'aol.ru', 'service.ru', 'supermail.ru', 'online.ru', 
               'webmail.ru', 'protonmail.ru', 'example.org', 'trashmail.ru', 'tutanota.ru', 'quickmail.ru', 'hotmail.ru', 'fastmail.ru', 
               'outlook.ru', 'inbox.net', 'mailhub.ru', 'example.net', 'mymail.ru', 'newmail.net', 'live.net', 'epost.net', 'myemail.net', 
               'coolmail.net', 'customdomain.net', 'icloud.net', 'smtp.net', 'msn.net', 'fakemail.net', 'yourmail.net', 'message.net', 
               'mail.ru', 'digitalmail.net', 'custommail.net', 'zoho.net', 'yandex.net', 'directmail.net', 'gmx.net', 'inbox.net', 
               'yahoo.net', 'mailservice.net', 'message.net', 'example.net', 'ezyemail.net', 'randommail.net', 'mxmail.net', 'gmail.net', 
               'mailbox.net', 'mailinator.net', 'mail.net', 'hushmail.net', 'testmail.net', 'aol.net', 'service.net', 'supermail.net', 
               'online.net', 'webmail.net', 'protonmail.net', 'example.org', 'trashmail.net', 'tutanota.net', 'quickmail.net', 'hotmail.net', 
               'fastmail.net', 'outlook.net', 'inbox.net', 'mailhub.net', 'example.net', 'mymail.net']
    
    for _ in range(n_rows):
        if data_target == "email":
            email = generator_function()
            local_part = email[: email.find("@")]
            domain = random.choice(domains)
            generated_value = f"{local_part}@{domain}"
            return_list.append(generated_value)
        
        elif data_target in ["date", "past_datetime", "time", "past_date", "date_time"]:
            generated_value = generator_function()
            
            if isinstance(generated_value, datetime.datetime):
                # Если дата без временной зоны, локализуем ее
                if generated_value.tzinfo is None:
                    aware_datetime = tz.localize(generated_value)
                else:
                    aware_datetime = generated_value.astimezone(tz)
                return_list.append(aware_datetime.isoformat())
            else:
                # Если сгенерированное значение — не datetime, возвращаем как есть
                return_list.append(str(generated_value))
        else:
            return_list.append(generator_function())
    return return_list


def check_type(value, planned_type=str):
    """This method checks if the handed value's type is equal to planned.
    
    :param value: The value to check
    :type value: Python object
    :param plannd_type: The intended python object type
    :type planned_type: Python type
    :return: None
    :raises ValueError: If types are not equal
    """

    #error_message = "Value of {} was expected to contain {} but was {}."
    #if type(value) != str:
    #        raise ValueError(error_message.format(
    #            namestr(value),
    #            "str",
    #            str(type(value))
    #        ))
    pass


def check_value_is_not_less_than(value, compare_value):
    """This method checks if a value is larger than a certain value.
    
    :param value: Object to check
    :param compare_value: Threshold to compare with
    """
    if value < compare_value:
        raise ValueError("n_rows must be at least {}, but was {}.".format(
            str(compare_value),
            str(value)
        ))

def check_value_is_not_more_than(value, compare_value):
    """This method checks if a value is smaller than a certain value.
    
    :param value: Object to check
    :param compare_value: Threshold to compare with
    """
    if value > compare_value:
        raise ValueError("n_rows must be at least {}, but was {}.".format(
            str(compare_value),
            str(value)
        ))



def namestr(object):
    """This method returns the string name of a python object.
    
    :param object: The object that should be used
    :type object: Python object
    :returns: Object name as string
    """
    for n,v in globals().items():
        if v == object:
            return n
    return None


def list_to_string(value_list):
    """Method concatenates a list of values to comma seperated string
    
    :param list: Python list of values
    :type list: List of strings
    :returns: Comma separated string
    :raises ValueError: If value_list is empty
    :raises ValueError: If value_list entries are not type str
    """

    # check if list contains elements
    if len(value_list) == 0:
        raise ValueError(error_02)

    # check if all values are type str
    for element in value_list:
        check_type(element, str)

    # return result
    return ", ".join(value_list)
