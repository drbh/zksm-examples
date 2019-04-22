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
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/setup", func(c *gin.Context) {

		var (
			s []int64
		)

		s = make([]int64, 10)
		s[0] = 12
		s[1] = 42
		s[2] = 61
		s[3] = 71
		s[4] = 1230
		s[5] = 121
		s[6] = 99421
		s[7] = 4210
		s[8] = 7111
		s[9] = 4214

		p, _ := zkproofs.SetupSet(s)

		spew.Dump(p.Signatures)

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

	r.GET("/prove", func(c *gin.Context) {
		var r *big.Int
		r, _ = rand.Int(rand.Reader, bn256.Order)
		// spew.Dump(r)

		var pm = new(zkproofs.ParamsSet)
		var pubke = new(bn256.G1)
		var Hke = new(bn256.G2)
		jsonMap := make(map[string]string)

		var snutr = "eyIxMiI6Ikt1a1M5dmJZenJQd3dnNUJIT1VRQURHOXVjc2U4c0tNTVBiQ1RCV1BlS0lLTlg3Sm1jRmk3T1BUWS9Kb1BxeUNEbjhjbWVuelFvVkNDR05nQjhVdVFCQ0g2L3Zqaklxd2pXMUpzZnNCWWRrSmxOK0xLalBQcG9RQndMVFFQSEllQnIxSHVuVlZsSUM4Q0J3RUxxSmoyL2dqcUJMblFqRTBtZzI3SXNrQTRiaz0iLCIxMjEiOiJGZXZhQk1FTlQzODMzT0x4RzVTT2JJVURNc293QlRHS3YyTUxxUTNwTEg0ZWZQZHZvcndid3JWSGp0dFhidVVjTTY2NlNBczVpN1dzQmlpaFdmNmxKdzVycTI4SUZoZFkzZnlGS0p2aXl3ZnZlU0gzR0FjaDV0UDlYci9VQXNUVEovZUtCdUlLVVhVUGxRMDN6eEExMUFUNmdLQVJrS3J0QlNNWTNMOVl6Qkk9IiwiMTIzMCI6IktoeWY1YWNoRms0cFVJMGdidnFRNDlRMUpZcGVLM3liNmRRbFRIdkJmcVVrRTNjSFdFM0lYRndRZGhIME04NE83aDVkQmRRK2VrVVFWQ2w5Njh2ZVF4MVc4bHlYMzRseUpkalZXOWJKeGJtRGhLNnQ5RzdqMWZoYmk2SkVZTW9DSHpQeTB5dnFINWM0WFVsZlY3aHpEdTBHK0I4TDRwTUZIWGp0WUwyNWNrUT0iLCI0MiI6IkV5UVovZTBBWnhHdkZLVUpHb0VaNGZFVTdLSlRhbjVvOGE0Z1JGWFBIdzRhOU41LzZPbzhKVmN0QUJyZ0xSWWEwZUxzN3JZM2xBek1nVFRzcFd3N0FTOTE1am4wazl2RnY0d2xOYk5nb3psWncydTFoZjJkNi95RU1OdjE3dXB0TFRtK2JOMDNuUWFhUzZjeDE3V0lvdlVYYWI5NURQVjB0S2RnMWFnUHJ0Zz0iLCI0MjEwIjoiSWVGK0g3SDZSak1MMFJrQTFBNVE0NXI2T3g4bFNCSW5QVExWUkNlTHVwVVNObHdOcm54RUQ3ODBFZE00TVhDa2g1eCtzTUYyQ0lxUjQrM01CVmlhQlNiZWFPbGxya0haVk9xWXB6VGQ3L3ErUnpjSEcyZFc4RCszL25DYWNIZGNKM1QwSnAxTnFnZU1zSHdUOE5yL21PcWc1MjVxcGJDa20zR0FLWTZlcyswPSIsIjQyMTQiOiJKSnJXaTVFS3gxL3Y2aXdVbzBwY1YxQUhVdU9RY3V2Y2liUXp1T0ZsWkRRYXFMMjQrM2Q0RDlHV25HQ0paSTlZYlZiWDV5S05SQWtPZjFMalJkMlEzaUR3cTJrR2YyalFGd0JuUzFjanhrTnZtbEtsNERpb2U5TDNrL2pHRzBwL0lYbVFQQVE4SWxaWjdZTEp4VlJYY0lPc3VrSVlyc1EvZ1JOMTY4RWpyZFU9IiwiNjEiOiJGTDlKV1ZQQzJGVkJ2WnVCeWU0OHIwRW9VVGtSb1E3Vld4RFlldU51a1EwZkNpRjVrNFF6Y1VvNkJBaVNWVzVFL3diUDNudVBFaHlWM2NFeFFIK0l2eEZEeFNDemxCOTU2SkJwdnVYdnpuWmxHd25od28wVEhNa09laDh0aElqaUthazRuNk80TE5RSjFmZjR0NFRhTWhvMHJnTUdMd0dBeFRwekNNdXpDOTA9IiwiNzEiOiJBamdyVkNnTDZ2N1l4S052MXVWVGhlcTZHdzNkaUlxSVRJUG96ODk0dG5JS2FLTjlIYnVNSkU4eFJVbGM2M05IelJjU0s1TXR0WmpRdnMvYm55TEc0d3gwVDNybE82cFBlS1E4L25VVlRZL3VRMGpUTi83N0Rlb3AxbmVwODVOSUgyckwyWlZObjJrYitrdVd6cnk2bWNnTnZ4dElJeCtESkZQVnd1SXUrWEE9IiwiNzExMSI6IkFBR29XN2lidEk3ejFoazRWMUMyTzBLUkg2OUJEb2tRV2FPV0pSekZIM2dlTW5pWFVjcGRSWkxoMDJWMEt0OXdRbWFuczZJZ0JsUy84Wm8waDQwamZRM0hZR2R3amV6SVZBL1lqQWNBTEZNZlZBSWdpeTgvL0oyT2NLZzNqdmNKS1g3Yyt6TXk0RklLM2ZKUXFmbTduOVdiWVZNR0p4eVVqV1VtRmdNOHh5Zz0iLCI5OTQyMSI6IkZMemdqTWI2YTZ5MXBTNVczb1FqcTg1NmZPektqL08vV0dtUnhlZG1GTjBEeURMZzdTRFhlNVhFKzVQQ0UyQ2xNaHlSRkdWZFg5MEZEUThyY0RKS01SZXo5UzdpQlVTdUY5SXlLZG9Pek1Lbk10ME9oRXh3YXdoUnM4VXdTQW5sRXI0T3VjOThyRzZUVEJ6N3NtT29SblJPelJkZUxJT3c1YXZnbXRUU0VTdz0ifQ=="
		sSigs, _ := b64.StdEncoding.DecodeString(snutr)

		json.Unmarshal(sSigs, &jsonMap)
		newsigs := make(map[int64]*bn256.G2)

		for k, v := range jsonMap {
			var myk = new(bn256.G2)
			var iky, _ = strconv.ParseInt(k, 10, 64)
			sSigs, _ := b64.StdEncoding.DecodeString(v)
			newsigs[iky], _ = myk.Unmarshal(sSigs)
		}

		var pub = "KBHHMtSNyJSufc0KD67kKAkvc57FB2I0WdVIYVgMrXwUkDRUJnqM6Cd/EnFq2IxLTphyKphfI/WXsa1J+XmQCA=="
		sDec, _ := b64.StdEncoding.DecodeString(pub)
		var val, _ = pubke.Unmarshal([]byte(sDec))

		var h = "IGBm3pCN+GdujqYteTWGAddoA2ok03QqELPGC4RsMlgtC4SMq+329hGZOODHKduGCsSpxL2G4hAsACLI9WJ0zhNf9r6WbuauLMEHZc+DbTpBHCHK2T4dbiujnrHVxTKML9ni9PJnTBanSQM2ggax+H1+8oM2qiVmU1AB0ufhbdE="
		sH, _ := b64.StdEncoding.DecodeString(h)
		var Hval, _ = Hke.Unmarshal([]byte(sH))

		pm.Kp.Pubk = val
		pm.H = Hval
		pm.Signatures = newsigs

		proof_out, _ := zkproofs.ProveSet(12, r, *pm)

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

	r.GET("/verify", func(c *gin.Context) {

		var pm = new(zkproofs.ParamsSet)

		var pubke = new(bn256.G1)
		var pb = "KBHHMtSNyJSufc0KD67kKAkvc57FB2I0WdVIYVgMrXwUkDRUJnqM6Cd/EnFq2IxLTphyKphfI/WXsa1J+XmQCA=="
		sDec, _ := b64.StdEncoding.DecodeString(pb)
		var val, _ = pubke.Unmarshal([]byte(sDec))
		pm.Kp.Pubk = val

		var Hke = new(bn256.G2)
		var h = "IGBm3pCN+GdujqYteTWGAddoA2ok03QqELPGC4RsMlgtC4SMq+329hGZOODHKduGCsSpxL2G4hAsACLI9WJ0zhNf9r6WbuauLMEHZc+DbTpBHCHK2T4dbiujnrHVxTKML9ni9PJnTBanSQM2ggax+H1+8oM2qiVmU1AB0ufhbdE="
		sH, _ := b64.StdEncoding.DecodeString(h)
		var Hval, _ = Hke.Unmarshal([]byte(sH))
		pm.H = Hval

		var proof_out = new(zkproofs.ProofSet)

		var new_A = new(bn256.GT)
		var str = "AFOOly8St0FrGbLwYuW2ELV6eLwd79U/+giH9IeSoWAEZWlVgyqRF13l+yI+SEvd8tuqtBlHR+ltn/M5x5c5TxREfl8sX6ueBbpdblYzKibbsQynGdWZ0+I3M7v0+gZnGVYTNk1YzMUSYfUZT69KmoJIE7OyAf2/aYhTG9HIs7ISjeD0lqiLjFqHmcA45CztC1ckPK/rhaeSwvrFl+yjzxdR6B4MFu9hKVjCOzYj6qr/w6P8QEO/m2/IFhY1wVyAGBajD8aYpTUZS3jhWPSOI6K/IX5QKGKYQ8NFljHm1pwHDgLFiHZcg7z0OdeoWAZkjGEUmct5OlHucoHo56D1lhFvJDLsCuv4qD5oGEtvkgM/XxD+5uaCdSUiOKUUS6vsLutUrAfAGnnkIr3rLLaUM7QmadAzIlbVPuruhVZUd1MRWhpPOAekOBgBOjje8VfTXaqevjTmR/GV4gAaJ3ADLQo3WLySyjGzj5F7ukAdiisE/rUW1O3UipvlEEI2WjOb"
		sstr, _ := b64.StdEncoding.DecodeString(str)
		var data, _ = new_A.Unmarshal([]byte(sstr))
		proof_out.A = data

		var new_V = new(bn256.G2)
		var strV = "Lx/yeXlHJSRGDtJ1z+GibqU1Wom090k8zHYH4v6abyoFKF/pfzXrq6kCPouQf+vq6YjgBeP5G5hx0BbOVvUzwAUAdgQYibFnQg/liHDm6oc8mXaalxASNKyLZWyH4yrhA3QY2tItBwHDnMgVF3l4qbxi9+mSuunf7koQF/s5HlA="
		sstrV, _ := b64.StdEncoding.DecodeString(strV)
		var dataV, _ = new_V.Unmarshal([]byte(sstrV))
		proof_out.V = dataV

		var new_D = new(bn256.G2)
		var strD = "K8P6IJPQAUDGfDDrfRU+vcudhStqwYQmIEnqk5fv7UoKlUiXG1mBqIOP2BVEo6zLiU6Rsx00MpdrkwCZHOE/giZXyVXFzxSIzGriFZOa2qHC7sXfcE4uDfvSO+RrKqiSIr2Ubgc/tZO7vQsOYxAaa4/idPZ4S3cfJzaE3LVrYbM="
		sstrD, _ := b64.StdEncoding.DecodeString(strD)
		var dataD, _ = new_D.Unmarshal([]byte(sstrD))
		proof_out.D = dataD

		var new_C = new(bn256.G2)
		var strC = "HAQkk7QxNikOsXN4yprss9ZMssDTjo5cg7yy8JYw0e0DD/IrPNbE4hQh+WLC3x0LTSTrfnKcL3F9c4T7ljVbdBuFBNNXT6x+FacRtqDSUb1TrmTK49SsioF6TGJz9z8VJ5237W/fwmmLqRnZpFQYk1JqgwoGJqcE6QMCuxiJfDg="
		sstrC, _ := b64.StdEncoding.DecodeString(strC)
		var dataC, _ = new_C.Unmarshal([]byte(sstrC))
		proof_out.C = dataC

		proof_out.Zr = zkproofs.GetBigInt("3116264101431494735950074624405961368626094989987397223809018573285082624904")
		proof_out.Zsig = zkproofs.GetBigInt("12633297896622030266045405375302026910081702495084548474562588435268755564119")
		proof_out.Zv = zkproofs.GetBigInt("14966269820474361960519298744011815110385788368318562302255605942208112802693")
		proof_out.Cc = zkproofs.GetBigInt("594934274205934811641591747988179780692057329804455304870096990666651305935")

		spew.Dump(proof_out)
		result, _ := zkproofs.VerifySet(proof_out, pm)

		spew.Dump(result)

		c.JSON(200, gin.H{
			"message": "res",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
