package main

import (
	"crypto/md5"
	"fmt"

	"github.com/satori/go.uuid"
)

func changeAcronisGuidEndian(id uuid.UUID) uuid.UUID {
	res := id
	res[0], res[1], res[2], res[3], res[4], res[5], res[6], res[7] =
		res[3], res[2], res[1], res[0], res[5], res[4], res[7], res[6]
	return res
}

func main() {
	centralizedPlanID, _ := uuid.FromString("af70a997-ca45-43d0-b478-afc6b7e5858f")
	//localPlanID, _ := uuid.FromString("954A2960-11C3-0D4D-1342-FA71A73CA90B")
	centralizedPlanID = changeUuidEndian(centralizedPlanID)
	instanceID, _ := uuid.FromString("2ECB7D51-3467-4F11-A24C-BC04BFE83395")
	instanceID = changeUuidEndian(instanceID)
	// var tempID uuid.UUID
	// for i := range planID {
	// 	tempID[i] = planID[i] ^ instanceID[i]
	// }
	// tempID[0], tempID[1], tempID[2], tempID[3], tempID[4], tempID[5], tempID[6], tempID[7] =
	// 	tempID[3], tempID[2], tempID[1], tempID[0], tempID[5], tempID[4], tempID[7], tempID[6]
	hash := md5.New()
	hash.Write(instanceID[:])
	hash.Write(centralizedPlanID[:])
	protectionID, _ := uuid.FromBytes(hash.Sum([]byte{}))
	protectionID = changeUuidEndian(protectionID)
	fmt.Println(protectionID)
}
