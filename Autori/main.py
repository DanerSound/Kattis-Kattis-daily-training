# Stringa di input
input_string = input()

# Dividi la stringa in una lista di sottostringhe utilizzando il carattere "-" come delimitatore
substrings = input_string.split("-")

# Estrai la prima lettera da ciascuna sottostringa e convertila in maiuscolo
first_letters = [substring[0].upper() for substring in substrings]

# Unisci le lettere estratte in una nuova stringa
result_string = ''.join(first_letters)

# Stampa la nuova stringa risultante
print(result_string)
