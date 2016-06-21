
def combine_files(base_folder_name, source_file_name_list, dest_file_name):
    dest_file = open(base_folder_name + dest_file_name, 'w', encoding='utf-8')
    for source_file_name in source_file_name_list:
        print(source_file_name)
        source_file_full_name = base_folder_name + source_file_name
        source_file = open(source_file_full_name, 'r', encoding='utf-8')
        file_content = source_file.read()
        dest_file.write('\r\n')
        dest_file.write(file_content)
        source_file.close()

    dest_file.close()


base_folder_name = 'c:/MyGo/src/matrix/static/ace/js/'
source_file_name_list = (
    'jquery.validate.min.js',
    'jquery-ui.min.js',
    'jquery-ui.custom.min.js',
    'bootstrap.min.js',
    'jquery.bootstrap.min.js',
    'ace.min.js',
    'ace-elements.min.js',
    'ace-extra.min.js',
)
dest_file_name= 'all1.js'
combine_files(base_folder_name, source_file_name_list, dest_file_name)

base_folder_name = 'c:/MyGo/src/matrix/static/ace/js/'
source_file_name_list = (
    'dataTables/jquery.dataTables.min.js',
    'dataTables/jquery.dataTables.bootstrap.min.js',
    'dataTables/extensions/buttons/dataTables.buttons.min.js',
    'dataTables/extensions/buttons/buttons.colVis.min.js',
    'dataTables/extensions/select/dataTables.select.min.js',
    'date-time/moment.min.js',
    'jquery.dataTables.fixedHeader.js',
    'jquery.dataTables.fixedColumns.js',
    'extend.datetime.js',
    'jquery.numberformatter.js',
    'datetimepicker/bootstrap-datetimepicker.min.js',
    'datetimepicker/locales/bootstrap-datetimepicker.zh-CN.js',
    'bootstrap-multiselect.min.js',
    'fuelux/fuelux.spinner.min.js',
    # 'jquery.ba-resize.min.js', # 不知道为啥，这个加入就会出错
)
dest_file_name= 'all2.js'
combine_files(base_folder_name, source_file_name_list, dest_file_name)


base_folder_name = 'c:/MyGo/src/matrix/static/ace/css/'

source_file_name_list = (
    'bootstrap.min.css',
    'font-awesome.min.css',
    'ace.min.css',
    'ace-fonts.min.css',
    'bootstrap-datetimepicker.min.css',
    'bootstrap-multiselect.min.css',
)

dest_file_name= 'all.css'

combine_files(base_folder_name, source_file_name_list, dest_file_name)