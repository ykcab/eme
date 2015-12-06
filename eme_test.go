package eme

// Test other data lengths using self-generated test vectors

import (
	"strings"
	"encoding/hex"
	"testing"
)

// unhex - strip newlines and decode hex
func unhex(s string) []byte {
	b, err := hex.DecodeString(strings.Replace(s, "\n", "", -1))
	if err != nil {
		panic(err)
	}
	return b
}

func TestEnc16(t *testing.T) {
	var v testVec
	v.dir = directionEncrypt
	v.key = make([]byte, 32)   // all-zero
	v.tweak = make([]byte, 16) // all-zero
	v.in = make([]byte, 16)    // all-zero
	v.out = unhex("f1b9ce8ca15a4ba9fb476905434b9fd3")
	verifyTestVec(v, t)
}

func TestEnc2048(t *testing.T) {
	var v testVec
	v.dir = directionEncrypt
	v.key = make([]byte, 32)   // all-zero
	v.tweak = make([]byte, 16) // all-zero
	v.in = make([]byte, 2048)  // all-zero
	v.out = unhex(
`630500884158a8d41216aaa6351e92ea7ca540a949e73f1d6069afef372eae226d426d8f4dbbce5
05051fcd596e7b32cdd1e44a40bcce0887ebbc160f779dafbb68bc25a40040987cdb5aba19b956e1
1dd69d7ce74d1d8788903254217300f3afd9a92cea395a89fd842256d3919ac303c1d1a21ec44cb2
898acabb3d0e04e05845e9183a64b4148f9614f62f79fa971799c896c58cd2b285cb00e40720b5c1
18ef1c0dbf5e51f09cd10db672b9c22c53d4bf5dabf16cfcbc382e3209e7ba27955b2afa055924d2
155319205f964492e0e88172f4ffadca35cd3e694fd3d803f610edc1696113996749bdb2a2dfb8fb
d92f0ef723146f93d969d1de4cc1c19efa33785f135a1203005ba2208fcd6b9f0d0e942993671b0c
5baefe7ed0bbfe070a9e1d26eed37228ceda64b70c25a59fd040703acc385c6e4093a8e258984371
c2d811c01723b87c820a55470501b245d97a0c7564c44e1e31a89457c1125198d27ac2dec67f7237
c781f9f622e420da471af577fa6ad8014bb1dad52dc450e37ebeeb1f1b7ba3c3d5796c6646bd052f
e66cf61d22f6e58efcbf0328d3b6e1088603e54bfd2da241c970fec58c7d21fc2b2f1b4b79552a92
500a7e6848a22ec47ee362cef9787af3b26439470992c8d103997395aefeff10802673d4630009e4
e5ed26038d5f05c589e64046274d5ff133398e4bab08ca71c0ce93dbd0657785b150c9e418339e8e
47d02ae7322b5b55364dfe10dd02d54dc5f2dbe82e55bcf0503b30d709301fc7f63cd15e7e124518
9405b1364b6074c0e493f7beb273859864e95643c7d914279d95f220247618b7a5c896d604e9ce88
e0a3ab434cfe97cabc520b20e2fffee958972e82e98b952cbf3567be560a8385e520b636829252ca
e16ac882d1d3dde447a90af52be7967f357a36382c9672a659ed0ffae9ee7b969c8c788f982fbd52
4efb6b852f65be5f371032691e9ef6e5a217607cf990e5351b251f841834045d9bd54633e8bb0d19
569377f805857a503a2642dfc06e2063dfd465dcae83b1e362394cdf20bd6a6d5dd0bacad7f2d407
e597b004e90683bfc23b92796bb0cdf67752765dbba2d1332848adcb83e60b7f3f8c397a3aa5de08
bf491676a18ee4a6ed9b342b33c081794365a72668f273a655e150fa36686aa5dfbc87ccda9cbe79
82e531044cd9fcea68ca1a5034ee64e0558b7c6af56808c0343506dccf39e4ca8634518e6420c107
b7d5fe89943e9f40044b2ae7f047d97275ec4763b41416fb3faff897e6ee0047f97b215b02ee09a3
f533cba2ee131a0aa68772860759325ea4dc7ab86528d0aff55f96d00a908ea32f9fa8bd0640d632
bf0d812f43beb03e6c6d4a654b41f27143fde4aa7ed16ae4bba939fbed8ae28dace83fcb49fd3ca9
10512902fc696419968e260901796f3af75d059c6e93a8f53b445f921e60d9429007c78107309a86
102b923dbb2f1fa04f84539c7c58ed45b40d4b267541d8f2f932ae75dfc13ef906089e1c847635aa
ada201e32301f326a733f4a99caa56fa00ccd6de2752a4495408bb4c930eaf84c97a28767a036f65
627efa868adeee1cf47840e67ac4489381da1513dd3d1707a5ad91ac4ba1c35714b1cd77e0555161
2cdfb88187a5a39afd5b2a54f88aeb9605c1e01cd7a976a168c6e40eddcb20f09fe54a07a7c971ed
a08789b41e450f46a30c2b3494f30a06073631d151ae048793fe32ff08a08b8693d1da71451f7ff2
438d7a521f8c97e4bc8ee744c1574d296c799d3f94e81522ce81def780660cee061d3110c1f18133
dca1daabb2268aad76327f3cb6ae8c96fc66b5506f5644f2652d022d7b616209970f78b8e5c9a74f
5d1fc31037d01733bff097a661d4331adc087b84e32afb9d82a0614ea459022918f8ca60e74b7c68
65d3e86e6ae4acf9546985dc5ccf242dd9ab0486bde99a1e656499cbfceef8de5beafda89965bf20
8c70f4a0f8f677792962b09e0a355d53bbacb4a7acc16e50aa483994d85ca7e28a34ee52843d4f27
acfb3b5347b89099fd01642b605f87f82283bb146aa319048462fe11a71d0fa064e825f69a502ae5
39240b13a2568383f80c30ebb9a29d2de4c45e6405bd8d1586fe5eef0477d3c21434b3471c0f864d
d4c2f135278645ef1b1ebb4d6954f92763d787ad3d0bb48409782cd803a877027ae105511cdaefe0
110f9f6c65f2a34ad12532f4715c0341346dc06bcc4bb0e5c04ffdfd3f625e7e3811a760a2cfa9cb
e1c2914882c192b793aeeb3ab8b03fd8be7e248fc9e5aa0c651d4bb5318c8a266b735ddc167d8374
99a3b38dcdfe188e25172eb974eb0eff20cc5c6e42e505fea7a3b6edfdf93b7062737ea00f72a415
c4dfd3cd4fab2dad812d88d02690f606b52dcfb0ab220dc65a7dee6f31e951324a1c442adfe9cf67
2646876bedc2c2cb71bac583f3cf4ea6769eb54d22bf22b0d908a297caa13767ed14c44b738799e8
21885b2b2988a9b88a2f35bf8d307543af9007fadcfcc273f8029c2fe3d59449c5f193ca7f8986ca
7451d9a3ee14335b10868421b64ea1605d48d4fc22e378f82b7fe3e88aada30b4c545408b87b2451
1e371f25d1a94e7b78a6a7aec848e6f82614f96ef5ed8705e0adec0f889c1860cf8211d716fddf31
9862290019a4387fb77c89a5ae7d0074fa72024f2e5096d39ebac8e8ed02191879585394987fea3a
07a720d6880910629cdcd83a02ffd98830d2f57ee6c53b5cfedb42ec7c3d925cf3080c343f21711a
90f117e9eaf91402d09b83e83bd18d1c4e3d165a8b9bac7adca12cb0147bfc4c6e2166c57b8182e6
3fbc698881ed5b329eb491eb98050a922c15804021013bf08942db9ee6d8f2a2c4eb93771340ed9e
323d09e4256b7e5ac`)
	verifyTestVec(v, t)
}
