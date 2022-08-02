def isValid(s: str) -> bool:
    symmap = []
    for c in s:
        if c == '{' or c == '(' or c == '[':
            symmap.append(c)
        else:
            lastsym = symmap.pop()
            if lastsym == '(' and c != ')':
                return False
            if lastsym == '{' and c != '}':
                return False
            if lastsym == '[' and c != ']':
                return False

    if len(symmap) == 0:
        return True
    else:
        return False

print(isValid(s = "()"))

