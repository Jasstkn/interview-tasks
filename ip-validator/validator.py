def has_leading_zeros(input):
    return input[0] == '0' and len(input) != 1

def ip_validator(ip):
    input_octets = ip.split(".")

    if len(input_octets) != 4:
        return False

    for octet in input_octets:
        if has_leading_zeros(octet):
            return False

        try:
            octenInt = int(octet)
        except:
            return False

        if octenInt < 0 or octenInt > 255:
            return False

    return True

