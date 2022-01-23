package entity

type Sg_server_config struct {
	Identity                 int32                    `json:"identity"`
	Name                     string                   `json:"name"`
	Public_domain_game       string                   `json:"public_domain_game"`
	Public_domain_platform   string                   `json:"public_domain_platform"`
	Internal_domain_game     string                   `json:"internal_domain_game"`
	Internal_domain_platform string                   `json:"internal_domain_platform"`
	Internal_k8s_svc         string                   `json:"internal_k8s_svc"`
	Max_user                 int32                    `json:"max_user"`
	Status                   int32                    `json:"status"`
	Remark                   string                   `json:"remark"`
	Rongyun                  Sg_server_rongyun_config `json:"rongyun"`
}

type Sg_server_rongyun_config struct {
	Chatroom_id_map []Sg_ChatroomIdMap `bson:"chatroom_id_map" json:"chatroom_id_map"`
	Gm_uid          string             `json:"gm_uid"`
	Keeplive_msg    string             `json:"keeplive_msg"`
}

type Sg_ChatroomIdMap struct {
	K int32  `bson:"k" json:"k"`
	V string `bson:"v" json:"v"`
}
