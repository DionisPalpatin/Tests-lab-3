import os
import random

sql_file_src = os.path.join(os.getcwd(), "ins.sql")
sql_file_res = os.path.join(os.getcwd(), "res.sql")
sql_file_res1 = os.path.join(os.getcwd(), "res1.sql")
sql_file_res2 = os.path.join(os.getcwd(), "res2.sql")

# data = None
# with open(sql_file_src, "r", encoding="utf-8") as f:
#     data = f.readlines()
    
# new_data = [data[0]]
    
# for el in data:
#     if not el.startswith("INSERT"):
#         new_line = "\t" + "(" + el[1 :]
#         ind = new_line.rfind(";")
#         new_line = new_line[:ind] + "," + new_line[ind + 1:]
#         new_data.append(new_line)
        
# with open(sql_file_res, "w", encoding="utf-8") as f:
#     data = f.writelines(new_data)        

# data = None      
# with open(sql_file_res1, "r", encoding="utf-8") as f:
#     seenl = list()
#     cl = 0
#     seenp = list()
#     cp = 0
    
#     data = f.readlines()
    
#     for i in range(1, len(data) - 1):
#         data[i] = data[i].split(", ")
#     data[-1] = data[-1].split(", ")
    
#     set_passwords, set_logins = set(), set()
    
    # i = 1
    # while i < len(data):
    #     if data[i][2] in seenl:
    #         dog_ind = data[i][2].find("@")
    #         data[i][2] = data[i][2][: dog_ind] + data[i][3][1 : len(data[i][3]) - 1] + data[i][2][dog_ind :]
    #     else:
    #         seenl.append(data[i][2])
    #         i += 1
    
    # for i in range(1, len(data)):  
    #     if data[i][2] in seenl:
    #         # print("here2:", i, data[i][3])
    #         # break
    #         cl += 1
    #     else:
    #         seenl.append(data[i][2]) 
          
    # for i in range(1, len(data)):  
    #     if data[i][3] in seenp:
    #         # print("here2:", i, data[i][3])
    #         # break
    #         cp += 1
    #     else:
    #         seenp.append(data[i][3])        
            
    # print(cl, cp)
    
    
# with open(sql_file_res1, "w", encoding="utf-8") as f:
#     for i in range(1, len(data)):
#         data[i] = ", ".join(data[i])
    
#     f.writelines(data)



# domains = [
#     'newmail.com', 'live.com', 'epost.com', 'myemail.com', 'coolmail.com', 'customdomain.com', 'icloud.com', 'smtp.com',
#     'msn.com', 'fakemail.com', 'yourmail.com', 'message.com', 'mail.ru', 'digitalmail.com', 'custommail.com', 'zoho.com',
#     'yandex.com', 'directmail.com', 'gmx.com', 'inbox.com', 'yahoo.com', 'mailservice.com', 'message.net', 'example.com',
#     'ezyemail.com', 'randommail.com', 'mxmail.com', 'gmail.com', 'mailbox.com', 'mailinator.com', 'mail.com', 'hushmail.com',
#     'testmail.com', 'aol.com', 'service.com', 'supermail.com', 'online.com', 'webmail.com', 'protonmail.com', 'example.org',
#     'trashmail.com', 'tutanota.com', 'quickmail.com', 'hotmail.com', 'fastmail.com', 'outlook.com', 'inbox.net', 'mailhub.com',
#     'example.net', 'mymail.com'
# ]


# domains_ru = list()
# for i in range(len(domains)):
#     domains_ru.append(domains[i].replace(".com", ".ru"))
    
    
# domains_net = list()
# for i in range(len(domains)):
#     domains_net.append(domains[i].replace(".com", ".net"))
    
    
# domains = domains + domains_ru + domains_net
# print(domains)


data = None
with open(sql_file_res1, "r", encoding="utf-8") as f:
    data = f.readlines()
    
    for i in range(1, len(data)):
        tmp = data[i]
        tmp = tmp.split(", ")
        tmp[-1] = str(random.randint(0, 2)) + tmp[-1][1 :]
        data[i] = ", ".join(tmp)
        

with open(sql_file_res2, "w", encoding="utf-8") as f:
    f.writelines(data)