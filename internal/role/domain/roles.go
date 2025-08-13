package domain

type Role string

const (
	RolePrimary                        Role = "Primary"
	RolePastor                         Role = "Pastor"
	RoleMediaMinistryLeader            Role = "Media Ministry Leader"
	RoleMediaMinistryMember            Role = "Media Ministry Member"
	RoleProductionMinistryLeader       Role = "Production Ministry Leader"
	RoleProductionMinistryMember       Role = "Production Ministry Member"
	RoleChildrensMinistryLeader        Role = "Children's Ministry Leader"
	RoleChildrensMinistryMember        Role = "Children's Ministry Member"
	RoleMusicMinistryLeader            Role = "Music Ministry Leader"
	RoleMusicMinistryMember            Role = "Music Ministry Member"
	RoleCreativeArtsMinistryLeader     Role = "Creative Arts Ministry Leader"
	RoleCreativeArtsMinistryMember     Role = "Creative Arts Ministry Member"
	RoleUsheringSecurityMinistryLeader Role = "Ushering & Security Ministry Leader"
	RoleUsheringSecurityMinistryMember Role = "Ushering & Security Ministry Member"
	RoleDisciple                       Role = "Disciple"
	RoleLeader                         Role = "Leader"
	RoleChurchMember                   Role = "Church Member"
	RoleEditor                         Role = "Editor"
	RolePhotographer                   Role = "Photographer"
	RoleSundaySchoolTeacher            Role = "Sunday School Teacher"
	RoleAdmin                          Role = "Admin"
	RoleMonitor                        Role = "Monitor"
	RoleLivestreamer                   Role = "Livestreamer"
	RoleITMinistry                     Role = "IT Ministry"
)

type RoleApplication struct {
	UserID           string
	IsLeader         bool
	IsPrimary        bool
	IsPastor         bool
	IsMinistryLeader bool
	MinistryID       *string // Optional, for ministry leader role
}

// RoleToLeaderMap is a map of leader role => member role
var RoleToLeaderMap = map[Role]Role{
	RoleChildrensMinistryLeader:        RoleChildrensMinistryMember,
	RoleCreativeArtsMinistryLeader:     RoleCreativeArtsMinistryMember,
	RoleMediaMinistryLeader:            RoleMediaMinistryMember,
	RoleMusicMinistryLeader:            RoleMusicMinistryMember,
	RoleProductionMinistryLeader:       RoleProductionMinistryMember,
	RoleUsheringSecurityMinistryLeader: RoleUsheringSecurityMinistryMember,
	RolePastor:                         RolePrimary,
}
