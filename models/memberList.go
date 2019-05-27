package models

import "fmt"

var DefaultMemberList MemberList

// Memberlist is interface of member list.
type MemberList interface {
	Add(member *Member) error
	Find(id string) (*Member, bool)
}

// InMemoryMemberList is a list of members in memory.
type InMemoryMemberList struct {
	members map[string]*Member
}

// NewInMemoryMemberList returns an empty InMemoryMemberList.
func NewInMemoryMemberList() *InMemoryMemberList {
	var memberList InMemoryMemberList
	memberList.members = make(map[string]*Member)
	return &memberList
}

// Add adds the given Member in the InMemoryMemberList.
func (m *InMemoryMemberList) Add(member *Member) error {
	if m.members[member.ID] != nil {
		return fmt.Errorf("Duplicated ID")
	}

	m.members[member.ID] = member
	return nil
}

// Find returns the Member with the given id in the Memberlist and a boolean
// indicating if the id was found.
func (m *InMemoryMemberList) Find(id string) (*Member, bool) {
	foundMember, ok := m.members[id]

	if ok == false {
		return nil, false
	}

	return foundMember, true
}

func init() {
	DefaultMemberList = NewInMemoryMemberList()
}
