SELECT * FROM auth_user WHERE login_name LIKE '%%s%'

# query type: (get find exec)
# find

# columns:
# id core.NullString
# login_name
# nick_name
# password
# phone
# weixin_id
# sex
# create_time
# update_time
# version

# filter:
# 100 ? engineer_id int64
# 100,200,300 %s type_id_list []int64

# '' datetime


id
login_name
nick_name
password
phone
weixin_id
sex
create_time
update_time
version