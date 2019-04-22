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

    # PERSON A -> PERSON B
    d = {
        "s": [12, 120, 1200, 12000, 24, 3]
    }
    resp = client.setup(d)
    # print(resp)

    # PERSON B
    h = resp["H"]
    pk = resp["Kp.Pubk"]
    p = {
        "val": 3,
        "h": h,
        "pub": pk,
        "sigs": resp["Signatures"],
    }
    resp = client.prove(p)
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
