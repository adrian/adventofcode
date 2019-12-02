def read_input(filename):
    with open(filename) as f:
        return [int(line) for line in f.readlines()]
