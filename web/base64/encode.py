import sys

# 文字列をバイナリに変換. 1文字8bit.
def str2bin(str):
    str2bin = str.encode()
    binary = ''
    for b in str2bin:
        binary += format(b, '08b')
    return binary

# 6bitずつに分割
def splitBin(bin):
    binList = [bin[i: i+6] for i in range(0,len(bin),6)]
    end = binList[-1]
    while 1:
        if len(end) < 6:
            end += '0'
        else:
            break
    binList[-1] = end
    return binList

# 6bitを変換表を基に文字に変換
def bin2str(bin):
    if int(bin,2) < int('011010',2):    ## A-Z
        return chr(int(bin,2)+65)
    elif int(bin,2) < int('110100',2):    ## a-z
        return chr(int(bin,2)+71)
    elif int(bin,2) < int('111110',2):    ## 0-9
        return chr(int(bin,2)-4)
    elif int(bin,2) == int('111110',2):
        return '+'
    else:
        return '/'

# 4文字ずつに分割
def splitStr(str):
    strList = [str[i:i+4] for i in range(0,len(str),4)]
    end = strList[-1]
    while 1:
        if len(end) < 4:
            end += '='
        else:
            break
    strList[-1] = end
    return strList

# main関数
def main():
    args = sys.argv
    if len(args) != 2:
        print(f'USAGE: python3 [FILENAME] [TEXT]')
        exit()
    binary = str2bin(args[1])
    base64List= ''.join([bin2str(i) for i in splitBin(binary)])
    base64 = ''.join(splitStr(base64List))
    print(f'binary: \n{binary}')
    print(f'EncodeBase64: \n{base64}')

if __name__ == '__main__':
    main()
