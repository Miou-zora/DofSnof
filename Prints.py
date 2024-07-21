from colorama import Fore, Style
import sys

def eprint(*args, **kwargs):
    """
    Print error in red color to stderr
    """
    print(Fore.RED, file=sys.stderr, end="")
    print(*args, **kwargs, file=sys.stderr)
    print(Style.RESET_ALL, file=sys.stderr, end="")
    sys.stderr.flush()

def wprint(*args, **kwargs):
    """
    Print warning in yellow color to stderr
    """
    print(Fore.YELLOW, end="")
    print(*args, **kwargs)
    print(Style.RESET_ALL, end="")
    sys.stdout.flush()

def iprint(*args, **kwargs):
    """
    Print info in blue color to stderr
    """
    print(Fore.BLUE, end="")
    print(*args, **kwargs)
    print(Style.RESET_ALL, end="")
    sys.stdout.flush()

def sprint(*args, **kwargs):
    """
    Print success in green color to stderr
    """
    print(Fore.GREEN, end="")
    print(*args, **kwargs)
    print(Style.RESET_ALL, end="")
    sys.stdout.flush()