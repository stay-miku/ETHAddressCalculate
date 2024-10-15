
# hex code
eth_code = "1234567890abcdef"

# base58 code
tron_code = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"


len = 8


print("eth:" + "|".join(["i{len}".replace("i", i).replace("len", str(len)) for i in eth_code]))
print()
print("tron" + "|".join(["i{len}".replace("i", i).replace("len", str(len)) for i in tron_code]))

