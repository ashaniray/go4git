package go4git

import (
	"bytes"
	"testing"
)

func TestReadPackIdx(t *testing.T) {

	rd := bytes.NewReader(PACK_IDX_DATA)

	idxs, err := GetAllPackedIndex(rd)

	if err != nil {
		t.Fatalf("Falied to read pack index - [%s]", err)
	}

	if got, want := len(idxs), 75; got != want {
		t.Errorf("object count = %d; want %d", got, want)
	}

	for i, idx := range idxs {
		if got, want := idx.Offset, packedIdxsWanted[i].Offset; got != want {
			t.Errorf("offset = %d; want %d", got, want)
		}

		if got, want := idx.CRC, packedIdxsWanted[i].CRC; got != want {
			t.Errorf("crc = %s; want %s", got, want)
		}

		if got, want := idx.HashAsString(), packedIdxsWanted[i].Hash; got != want {
			t.Errorf("hash = %s; want %s", got, want)
		}

	}
}

var packedIdxsWanted = []struct {
	Offset int
	Hash   string
	CRC    string
} {
	{6302, "04273da57c852a3c909a543bf6a1437139091ae1", "f4ba5890"},
	{6048, "0702f2027bbd7da762e06cf26097f6ebce7543b9", "78f1aa5b"},
	{3300, "09085414ee6ab45d5f28df2e6976f4aed941471e", "aebd6f1f"},
	{5401, "0c63f88031be6d322672024a6244bf9e7e00bc0c", "813da9b1"},
	{5692, "0f0aff4e1c99c62331519dc20bc5584eaa48dda3", "649888a7"},
	{3929, "131da4c508d554fcaed023a8afdeebe154f15444", "4ebe2a49"},
	{3855, "165fb5ee6a411a6f0312fb4642e13aec88766818", "8f294394"},
	{226 , "1697bf15fc565167eeca23d349824dd58a741cf7", "5fb9a7d4"},
	{5312, "17b726254f5990b0e3bbdad2501499c16a2aefa6", "2c9f5f25"},
	{6428, "182297434f3f967d320e15df029700a3fce12bcf", "ffcdb717"},
	{5919, "19b94cceab2e65909940c36d84a44163c66c9580", "4617494d"},
	{5117, "19bbecee167ac89444d16eedc597f7c8446bb78e", "37f855a8"},
	{5140, "1a516d4de657d145b6dc17f145a461ac7ed35987", "910df0fd"},
	{6664, "1d6e068e9598795a7fb5c8cfb3a1ed1b0811191f", "bce17dfe"},
	{4084, "1df71c7b057f872059074c23e8a189f92abfe14e", "4029f4f7"},
	{577 , "1ec1f6163004b93a8460f90bd03a5de9a90ea165", "faf9459b"},
	{4005, "2c5502eca91c73589d5fbd7328460858104458d3", "cd37b027"},
	{4810, "2c74b7fc8fe26bfcdfe17e3a95e1d1113eb075a6", "506bdf59"},
	{2945, "2f987200c2994abc86458a4bde3c7c3c89835600", "fa0536d8"},
	{7577, "33f9915107de019becb353807db0ddb4ceadf2b9", "6167e11e"},
	{5020, "34dfa56fbf7642b25ce69b1187f84775ddebf8d3", "a3a43dc8"},
	{4948, "3e15e1bdff93bff120ff1b6d7717d4a7350d4ca9", "2732cbe8"},
	{5775, "3e4838ed94cc1a93a17aa1175ab4495c76a00ec3", "7a269b35"},
	{6769, "3e6d58c2cc39920f7c1666aaf856ce20f6475c2f", "98929652"},
	{8085, "3e9cef68b1268edcb2f670bf2286c95564ac3661", "a9515b80"},
	{4590, "4274d5d2b2032991a9b4cf09f5e01a09d55e3d14", "9d04afdc"},
	{3703, "4a4cfa0f5b036cd69b150422dfa9895777ad9112", "b8a48004"},
	{5541, "50d095c4f543363d615b355f733b9fcc4d251421", "7fd4f384"},
	{7804, "517a2fd94c0c0da7ed13933ef303694f244c3c70", "1dca6d9b"},
	{8777, "5182ef3bb52970b4245accb9a60392018aedd0f2", "8dc32e2f"},
	{12  , "53023333cddb3bfacb78a30cbc0dabcd20108426", "63ccd83d"},
	{8529, "5eb6e2ecbf76aec16e9c4b37817f364406340e81", "da7fbeb4"},
	{6169, "622b9712895a3b040b18de6f05ff49dee75f606e", "13afa899"},
	{5211, "6f7fe8380b483f8beed8fd832c1a0a1b0677fe3e", "cf39764c"},
	{8901, "7308f6aec24e84da09f5b4f28c8310840633a7bf", "c41b8939"},
	{1001, "77873667494454ce7aa494391d1c6820ba3bbf60", "18e80a19"},
	{1425, "77d3f0dba02e6188278afa69ed8214c26371a265", "5ff70b98"},
	{2388, "79f959612c20dfed2fc5e42fe20221153938d40a", "8f9047e7"},
	{8375, "7a7f7aedf5d334393dec2ffe145112b693478bcc", "9c8657a3"},
	{7333, "7d0235d2f679c788e34039ed20379cc9757a88bf", "183dbac8"},
	{6538, "7da7cc9efbfa3913c0ac8875944aeea75cefd839", "dfa88c01"},
	{3628, "83c72baae8b3cf907ce043c836a3fa4e24aaee4a", "630d625f"},
	{5192, "9128e3b4d7fa29e8d6432b089fdd0f91309d802a", "286b3c7d"},
	{4522, "9377c1048a11565122361ab51eb1f674d2b85110", "58650c4e"},
	{8222, "939de2dd828de6ca8a1af9c89a41ff740f438c15", "f08c404a"},
	{7077, "95a4696a2a4e8cfc01a26d91fe26410053d9aac7", "36c526a7"},
	{4738, "9c211498d2054323cc846b129095e6837419781b", "229aad83"},
	{3114, "9d97f9964a9c744696d76c23711194586ce9a395", "d275bfb5"},
	{6959, "9eecb159f2e29ee0467ea9af2a26778a3f7351a0", "be8feb90"},
	{7931, "a27fed8338cbb568de4435d7044144dc523f4e6a", "20a8826e"},
	{7692, "a29bbc3bf73fc9b2fc110b7995b0299505806da9", "ebe86659"},
	{5098, "a7054f5bb479a95fffc679b9c98f025a0c03c12b", "bafca35d"},
	{1845, "a760b13e0179d27890bdce3f136e4516b1259390", "4d569a83"},
	{4878, "a7c57d716b50c9df60b800075234b4960c67ec4c", "90c1b410"},
	{1218, "a7fde8a1c710964baa35b5440bf4257f4744dcc4", "09aa09db"},
	{9152, "ae8e4729bbf8b5ef5d8ba89cbc77ee848d040c74", "a8bf3cc4"},
	{7204, "aee5f68530b5ec5285b62797d7add2787303b2ce", "8e23f918"},
	{3348, "b027499ba2c7eb1412cd45a2dd7e569349761fee", "1d94bc3e"},
	{7464, "b384e756449cdf6634d081b12733b20a94fd79ce", "0665f882"},
	{2578, "b58544fb9700775ea53e013ea53b63d5ae3360dd", "ca8fd4de"},
	{6863, "b70f8bdefa50905363a2bda5ec85ddc5d8a7ccc2", "c38e2b23"},
	{2203, "bd4d6c1931d3f28560408f51e6be9f1a27e355ad", "4853e3da"},
	{1627, "c0033c6b4d0e6d36598f6cebbbc34405159f1364", "ac28b5d3"},
	{3780, "cac0cb235ce2a6cd80d427041fef6868cf983a10", "f678ce9c"},
	{5169, "cecde88d1b23e3a1cd5110c7bc2913952a8026bf", "b5e9d65d"},
	{792 , "d1973756e19294214b3ebff2f8896800fb994a4d", "612bbd35"},
	{2035, "d3717dcf242dd7265a767a401d0e9daca8b3564b", "31698b1b"},
	{5236, "d434f9043495aff2c6c5d0f47e64f6e64655bfd9", "362dd622"},
	{4163, "d64b8c84024c7dcd7534383f504f74fe3f50d1a5", "43a48920"},
	{8668, "e2dd8e4d57f29c1be299e295966b1330c4835ef4", "2850fad0"},
	{358 , "f7f57b5406e487399238304a1b353aef8e8b2494", "a1fcef59"},
	{2754, "f9845fb488fba75f325e34705d4b835131f5ee47", "38aa5feb"},
	{4238, "f9a8f3646ba25d2d9ae95a644590db11b324507a", "70b91cec"},
	{9050, "fab6ea6999c45fb4526e6ecd3a2d0955b3c1f421", "d6a820d2"},
	{4660, "fc02c57c9ce66586c4cdbddca28e0879995b40fd", "20c488a2"},
}
