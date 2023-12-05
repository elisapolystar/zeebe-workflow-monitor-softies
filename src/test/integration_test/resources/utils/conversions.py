import base64

def base64_to_bytes(base64_string):
    return base64.b64decode(base64_string)

def bytes_to_string(bytes):
    return bytes.decode('utf-8')
