import os
import filecontent 

def filefiller(file_extention):
    if file_extention == '.go': 
        return filecontent.golang_content()
    elif file_extention == ".py": 
        return filecontent.python_content()
    else: 
        print(" language not supported yet:created an empty file ")
        return filecontent.default_content()


def generate_file():
    dir_name = input("write the directory:")
    file_name = input("Enter the file name with extension (e.g. main.go):")

    # get file extention 
    extention = file_name[file_name.index('.'):len(file_name):1]

    # Crea la directory se non esiste
    if not os.path.exists(dir_name):
        os.makedirs(dir_name)

    # Componi il percorso completo del file all'interno della directory
    file_path = os.path.join(dir_name, file_name)

    with open(file_path, "w") as file:
        file.write(filefiller(extention))


generate_file()
