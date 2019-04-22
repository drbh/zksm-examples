package main

import (
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"

	// "github.com/drbh/zkproofs/go-ethereum/crypto/bn256"
	"github.com/drbh/zkproofs/go-ethereum/zkproofs"

	"github.com/gin-gonic/gin"
	"github.com/ing-bank/zkproofs/go-ethereum/crypto/bn256"
	"math/big"
	"os"
	"strconv"
)

type verifData struct {
	H   string `json:"H"`
	Pub string `json:"Pub"`

	A    string `json:"A"`
	C    string `json:"C"`
	Cc   string `json:"Cc"`
	D    string `json:"D"`
	M    string `json:"M"`
	V    string `json:"V"`
	Zr   string `json:"Zr"`
	Zsig string `json:"Zsig"`
	Zv   string `json:"Zv"`
}

type proverData struct {
	H    string `json:"h"`
	Pub  string `json:"pub"`
	Sigs string `json:"sigs"`
	Val  int64  `json:"val"`
}

type setupData struct {
	S []int64 `json:"s"`
}

func main() {
	r := gin.Default()

	r.POST("/setup", func(c *gin.Context) {
		var (
			s []int64
		)
		var sdata setupData
		c.BindJSON(&sdata)
		s = sdata.S
		p, _ := zkproofs.SetupSet(s)
		sigs := make(map[int64]interface{})
		for k, v := range p.Signatures {
			sigs[k] = v.Marshal()
		}
		jsonSigs, _ := json.Marshal(sigs)
		c.JSON(200, gin.H{
			"H":          p.H.Marshal(),
			"Signatures": jsonSigs,
			"Kp.Pubk":    p.Kp.Pubk.Marshal(),
		})
	})

	r.POST("/prove", func(c *gin.Context) {
		var proveDat proverData
		c.BindJSON(&proveDat)
		// spew.Dump(proveDat)
		var r *big.Int
		r, _ = rand.Int(rand.Reader, bn256.Order)
		// spew.Dump(r)
		var pm = new(zkproofs.ParamsSet)
		var pubke = new(bn256.G1)
		var Hke = new(bn256.G2)
		jsonMap := make(map[string]string)
		var snutr = proveDat.Sigs
		sSigs, _ := b64.StdEncoding.DecodeString(snutr)
		json.Unmarshal(sSigs, &jsonMap)
		newsigs := make(map[int64]*bn256.G2)
		for k, v := range jsonMap {
			var myk = new(bn256.G2)
			var iky, _ = strconv.ParseInt(k, 10, 64)
			sSigs, _ := b64.StdEncoding.DecodeString(v)
			newsigs[iky], _ = myk.Unmarshal(sSigs)
		}
		var pub = proveDat.Pub
		sDec, _ := b64.StdEncoding.DecodeString(pub)
		var val, _ = pubke.Unmarshal([]byte(sDec))
		var h = proveDat.H
		sH, _ := b64.StdEncoding.DecodeString(h)
		var Hval, _ = Hke.Unmarshal([]byte(sH))
		pm.Kp.Pubk = val
		pm.H = Hval
		pm.Signatures = newsigs
		var value = proveDat.Val // 12
		proof_out, _ := zkproofs.ProveSet(value, r, *pm)
		spew.Dump(proof_out)

		c.JSON(200, gin.H{
			"V":    proof_out.V.Marshal(),
			"D":    proof_out.D.Marshal(),
			"C":    proof_out.C.Marshal(),
			"A":    proof_out.A.Marshal(),
			"Zsig": proof_out.Zsig,
			"Zv":   proof_out.Zv,
			"Cc":   proof_out.Cc,
			"M":    proof_out.M,
			"Zr":   proof_out.Zr,
		})
	})

	r.POST("/verify", func(c *gin.Context) {
		var verifDataObject verifData
		c.BindJSON(&verifDataObject)
		var pm = new(zkproofs.ParamsSet)
		var pubke = new(bn256.G1)
		var pb = verifDataObject.Pub
		sDec, _ := b64.StdEncoding.DecodeString(pb)
		var val, _ = pubke.Unmarshal([]byte(sDec))
		pm.Kp.Pubk = val

		var Hke = new(bn256.G2)
		var h = verifDataObject.H
		sH, _ := b64.StdEncoding.DecodeString(h)
		var Hval, _ = Hke.Unmarshal([]byte(sH))
		pm.H = Hval

		var proof_out = new(zkproofs.ProofSet)

		var new_A = new(bn256.GT)
		var str = verifDataObject.A
		sstr, _ := b64.StdEncoding.DecodeString(str)
		var data, _ = new_A.Unmarshal([]byte(sstr))
		proof_out.A = data

		var new_V = new(bn256.G2)
		var strV = verifDataObject.V
		sstrV, _ := b64.StdEncoding.DecodeString(strV)
		var dataV, _ = new_V.Unmarshal([]byte(sstrV))
		proof_out.V = dataV

		var new_D = new(bn256.G2)
		var strD = verifDataObject.D
		sstrD, _ := b64.StdEncoding.DecodeString(strD)
		var dataD, _ = new_D.Unmarshal([]byte(sstrD))
		proof_out.D = dataD

		var new_C = new(bn256.G2)
		var strC = verifDataObject.C
		sstrC, _ := b64.StdEncoding.DecodeString(strC)
		var dataC, _ = new_C.Unmarshal([]byte(sstrC))
		proof_out.C = dataC

		proof_out.Zr = zkproofs.GetBigInt(verifDataObject.Zr)
		proof_out.Zsig = zkproofs.GetBigInt(verifDataObject.Zsig)
		proof_out.Zv = zkproofs.GetBigInt(verifDataObject.Zv)
		proof_out.Cc = zkproofs.GetBigInt(verifDataObject.Cc)

		spew.Dump(proof_out)
		result, _ := zkproofs.VerifySet(proof_out, pm)

		spew.Dump(result)

		c.JSON(200, gin.H{
			"message": result,
		})
	})

	argsWithoutProg := os.Args[1:]

	// spew.Dump(argsWithoutProg)
	var port = ":" + argsWithoutProg[0]
	spew.Dump(port)
	r.Run(port) // listen and serve on 0.0.0.0:8080
	// r.Run() // listen and serve on 0.0.0.0:8080
}
