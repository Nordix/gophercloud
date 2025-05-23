/*
Package keypairs provides the ability to manage key pairs as well as create
servers with a specified key pair.

Example to List Key Pairs

	allPages, err := keypairs.List(computeClient, nil).AllPages(context.TODO())
	if err != nil {
		panic(err)
	}

	allKeyPairs, err := keypairs.ExtractKeyPairs(allPages)
	if err != nil {
		panic(err)
	}

	for _, kp := range allKeyPairs {
		fmt.Printf("%+v\n", kp)
	}

Example to List Key Pairs using microversion 2.10 or greater

	client.Microversion = "2.10"

	listOpts := keypairs.ListOpts{
		UserID: "user-id",
	}

	allPages, err := keypairs.List(computeClient, listOpts).AllPages(context.TODO())
	if err != nil {
		panic(err)
	}

	allKeyPairs, err := keypairs.ExtractKeyPairs(allPages)
	if err != nil {
		panic(err)
	}

	for _, kp := range allKeyPairs {
		fmt.Printf("%+v\n", kp)
	}

Example to Create a Key Pair

	createOpts := keypairs.CreateOpts{
		Name: "keypair-name",
	}

	keypair, err := keypairs.Create(context.TODO(), computeClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", keypair)

Example to Import a Key Pair

	createOpts := keypairs.CreateOpts{
		Name:      "keypair-name",
		PublicKey: "public-key",
	}

	keypair, err := keypairs.Create(context.TODO(), computeClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Key Pair

	err := keypairs.Delete(context.TODO(), computeClient, "keypair-name", nil).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Delete a Key Pair owned by a certain user using microversion 2.10 or greater

	client.Microversion = "2.10"

	deleteOpts := keypairs.DeleteOpts{
		UserID: "user-id",
	}

	err := keypairs.Delete(context.TODO(), client, "keypair-name", deleteOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Create a Server With a Key Pair

	createOpts := servers.CreateOpts{
		Name:      "server_name",
		ImageRef:  "image-uuid",
		FlavorRef: "flavor-uuid",
		KeyName:   "keypair-name",
	}

	server, err := servers.Create(context.TODO(), computeClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Get a Key Pair owned by a certain user using microversion 2.10 or greater

	client.Microversion = "2.10"

	getOpts := keypairs.GetOpts{
		UserID: "user-id",
	}

	keypair, err := keypairs.Get(context.TODO(), computeClient, "keypair-name", getOpts).Extract()
	if err != nil {
		panic(err)
	}
*/
package keypairs
