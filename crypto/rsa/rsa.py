from Crypto.Util.number import *

e = 65537


"""
    公開鍵と秘密鍵作成
    公開鍵 [e, n]
    秘密鍵 [d, n]
    引数: p, q 素数
    戻り値： dict
"""
def makeKey(p, q):
    n = p * q
    phi = (p - 1) * (q - 1)
    d = pow(e, -1, phi)
    return {'public': e, 'secret': d, 'num': n}

"""
    暗号化関数
    引数: m 平文(bytes), (e, n) 公開鍵
"""
def encrypt(m, e, n):
    return pow(m, e, n)


"""
    復号化関数
    引数: c 暗号文(bytes), (d, n) 秘密鍵
"""
def decrypt(c, d, n):
    return pow(c, d, n)
