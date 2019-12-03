def read_input(filename):
    with open(filename) as f:
        return [int(line) for line in f.readlines()]

def read_file_as_int_array(filename):
    with open(filename) as f:
        lines = f.readlines()
        return [int(n) for n in lines[0].split(',')]
