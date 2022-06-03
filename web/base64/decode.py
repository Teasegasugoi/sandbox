import sys

UPPER_CASE = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
LOWER_CASE = 'abcdefghijklmnopqrstuvwxyz'
NUMBER_CASE = '0123456789'
OTHER_CASE = '/+'

def splitStr(str):
    return [str[i:i+4] for i in range(0,len(str),4)]

# 変換表を基にバイナリに変換
def str2bin(str):
    if str in UPPER_CASE:
        return format(ord(str) - 65, '06b')
    elif str in LOWER_CASE:
        return format(ord(str) - 71, '06b')
    elif str in NUMBER_CASE:
        return format(ord(str) + 4, '06b')
    elif str == '+':
        return '111110'
    else:
        return '111111'

def splitBin(bin):
    binList = [bin[i:i+8] for i in range(0, len(bin), 8)]
    if len(binList[-1]) != 6:
        return binList[:-1]
    else:
        return binList

def main():
    args = sys.argv
    if len(args) != 2:
        print(f'USAGE: python3 [FILENAME] [TEXT]')
        exit()
    binary = ''.join([str2bin(i) for i in args[1] if i != '='])
    binList = splitBin(binary)
    text = ''.join([chr(int(i,2)) for i in binList])
    print(f'binary: \n{binary}')
    print(f'DecodeBase64: \n{text}')

if __name__ == '__main__':
    main()