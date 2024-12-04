import os

def input_file(day: str, is_test: bool = False) -> str:
    parent_dir = os.path.dirname(os.path.realpath(__file__))

    if is_test:
        day = f"{day}_test"

    return f"{parent_dir}/../inputs/{day}"