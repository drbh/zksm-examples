import requests
import json


class ZKSMClient(object):

    def __init__(self, port):
        pass
        self.port = port  # 8080
        self.host = "http://localhost:"+str(self.port)+"/"

    def setup(self, inputData):
        url = self.host+"setup"
        payload = json.dumps(inputData)
        headers = {'content-type': 'application/json'}
        response = requests.request(
            "POST", url, data=payload, headers=headers)
        data = response.text
        return json.loads(data)

    def prove(self, inputData):
        url = self.host+"prove"
        payload = json.dumps(inputData)
        headers = {'content-type': 'application/json'}
        response = requests.request(
            "POST", url, data=payload, headers=headers)
        data = response.text
        return json.loads(data)

    def verify(self, inputData):
        url = self.host+"verify"
        payload = json.dumps(inputData)
        headers = {'content-type': 'application/json'}
        response = requests.request(
            "POST", url, data=payload, headers=headers)
        data = response.text
        return json.loads(data)


if __name__ == '__main__':

    client = ZKSMClient(8080)

    client2 = ZKSMClient(8081)

    # PERSON A -> PERSON B
    d = {
        "s": [7650, 948, 3598, 9086, 604, 1330,
              114, 158, 9902, 9267, 2943, 9991,
              3589, 3654, 3802, 1241, 2179, 3347,
              2330, 4860, 2690, 8761, 7050, 5258,
              7607, 3746, 9547, 9442, 4076, 2586,
              2806, 8666, 9959, 7172, 8659, 9698,
              2442, 4741, 6382, 4042, 4195, 4549,
              3448, 2335, 2077, 9216, 8202, 7328,
              6626, 2920, 9961, 657, 9014, 7745,
              3047, 9801, 6930, 3169, 6779, 3258,
              5143, 1045, 572, 1702, 6196, 3360,
              5625, 26, 7940, 1129, 307, 9044,
              4758, 6342, 6034, 4122, 148, 1622,
              5341, 4573, 2030, 8392, 1685, 9107,
              536, 4619, 8498, 5507, 2686, 4359,
              4565, 685, 9486, 8360, 9203, 8726,
              4680, 3880, 8956, 9864, 7878, 5425,
              2495, 1091, 8257, 6900, 927, 9515,
              8270, 711, 1301, 4242, 8880, 1738,
              3661, 4935, 6686, 9705, 5091, 6817,
              5575, 3183, 4426, 9752, 9280, 8534,
              3897, 7143, 6776, 9606, 1635, 6113,
              4013, 6592, 3148, 2865, 4812, 3387,
              7332, 8913, 2790, 770, 206, 5889,
              4716, 6964, 3095, 4368, 9557, 2497,
              5325, 3705, 2852, 9101, 4202, 2890,
              8673, 8465, 8611, 8352, 6252, 5855,
              8484, 1604, 9080, 1531, 485, 3265,
              9989, 9321, 7335, 8378, 545, 2388,
              2635, 6059, 5788, 818, 3944, 6632,
              1676, 4432, 5820, 1633, 1580, 3165,
              5936, 9216, 6355, 289, 5803, 7435,
              761, 6556, 6817, 1597, 1116, 5953,
              2067, 8912]
    }
    resp = client.setup(d)
    # print(resp)

    # PERSON B
    h = resp["H"]
    pk = resp["Kp.Pubk"]
    p = {
        "val": 8912,
        "h": h,
        "pub": pk,
        "sigs": resp["Signatures"],
    }
    resp = client2.prove(p)
    # print(resp)

    # PERSON B -> PERSON A
    v = {
        "H": h,
        "Pub": pk,

        "A": resp["A"],
        "C": resp["C"],
        "Cc": str(resp["Cc"]),
        "D": resp["D"],
        "M": str(resp["M"]),
        "V": resp["V"],
        "Zr": str(resp["Zr"]),
        "Zsig": str(resp["Zsig"]),
        "Zv": str(resp["Zv"]),
    }
    resp = client.verify(v)
    print(resp)
