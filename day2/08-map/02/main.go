/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-16 22:29:30
 * @LastEditors: 1372594487 1372594487@qq.com
 * @LastEditTime: 2025-12-17 01:09:38
 * @FilePath: \go-learning\day2\08-map\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"    // 用于格式化输出
	"sync"   // 用于同步原语
	"time"   // 用于时间处理
)

type UserSession struct {
	UserID  string
	LoginTime time.Time
	LastActivity time.Time
	IPAddress string
}

type SessionManager struct {
    sessions map[string]*UserSession
	mutex sync.RWMutex //解决并发安全问题
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*UserSession),
	}
}

func (sm *SessionManager) AddSession(token string ,UserID string, ipAddress string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	session := &UserSession{
		UserID: UserID,
		LoginTime: time.Now(),
		LastActivity: time.Now(),
		IPAddress: ipAddress,
	}
	sm.sessions[token] = session  
    
}

func(sm *SessionManager) GetSession(token string) (userSession *UserSession, exist bool){
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	session, ok := sm.sessions[token]
	if ok {
		session.LastActivity = time.Now() //更新最后活动时间
	}
	return session,ok
}
func (sm *SessionManager) DeleteSession(token string){
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	delete(sm.sessions,token)
}

func (sm *SessionManager) ClearExpiredSessions(timeout time.Duration){
		sm.mutex.Lock()
		defer sm.mutex.Unlock()

		for token, session := range sm.sessions {
			if time.Since(session.LastActivity) > timeout {
				delete(sm.sessions,token)
			}
	}
}

func main() {
    sessionManager := NewSessionManager()
    sessionManager.AddSession("token1", "user1", "192.168.1.1")
	
    if session, ok := sessionManager.GetSession("token1"); ok {
        fmt.Println("User ID:", session.UserID,"session",session)
    }
	sessionManager.ClearExpiredSessions(time.Hour)

	sessionManager.DeleteSession("token1")
}