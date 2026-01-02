/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-27 02:20:50
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-27 15:05:55
 * @Description: File description
 */
package user

import (
	"sync"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

type UserStore struct {
	users  map[int]User
	mutex  sync.RWMutex
	nextID int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *UserStore) AddUser(u *User) (User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user := User{
		ID:        s.nextID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: time.Now(),
		Active:    true,
	}
	s.users[s.nextID] = user
	s.nextID++
	return user, nil
}

func (s *UserStore) GetUser(id int) (User, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, ok := s.users[id]
	return user, ok

}

func (s *UserStore) GetAllUsers() []User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}
