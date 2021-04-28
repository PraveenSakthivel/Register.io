import jwt

struct = {"name":"ps931","classHistory":{"192:101":400,"192:102":400},"specialCases":{"10":True}}
encoded_jwt = jwt.encode(struct, "KZQVGKt0WglphtNyME8912pa9RlnSZ1s8Xcdqe5OnQ", algorithm="HS256")
print(encoded_jwt)