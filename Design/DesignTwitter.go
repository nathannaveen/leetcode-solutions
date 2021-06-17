package main

import "fmt"

type person struct {
	id    int
	posts []int
}

type follow struct {
	id int
	f  []int
}

type Twitter struct {
	allPeople []person
	m         map[int]int
	followers []follow
	m2        map[int]int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		allPeople: []person{},
		m:         map[int]int{},
		followers: []follow{},
		m2:        map[int]int{},
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	fmt.Println(this.followers)
	if this.m[userId] == 0 {
		this.allPeople = append(this.allPeople, person{id: userId})
	}
	this.allPeople[len(this.allPeople)-1].posts = append(this.allPeople[len(this.allPeople)-1].posts, tweetId)
	fmt.Println(this.allPeople[len(this.allPeople)-1].posts)
	if len(this.allPeople[len(this.allPeople)-1].posts) > 10 {
		this.allPeople[len(this.allPeople)-1].posts = this.allPeople[len(this.allPeople)-1].posts[1:]
	}

	for i, user := range this.followers {
		if user.id == userId {
			for _, follower := range user.f {
				if this.m[follower] == 0 {
					this.allPeople = append(this.allPeople, person{id: follower})
				}
				this.allPeople[i].posts = append(this.allPeople[i].posts, tweetId)
				if len(this.allPeople[i].posts) > 10 {
					this.allPeople[i].posts = this.allPeople[i].posts[1:]
				}
				fmt.Println(this.allPeople[i].posts)
			}
		}
	}

	this.m[userId]++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by
users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	return this.allPeople[len(this.allPeople)-1].posts
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.m2[followerId] == 0 {
		this.followers = append(this.followers, follow{id: followerId})
	}
	this.followers[len(this.followers)-1].f = append(this.followers[len(this.followers)-1].f, followeeId)
	this.m2[followerId]++

	if this.m2[followeeId] == 0 {
		this.followers = append(this.followers, follow{id: followeeId})
	}
	this.followers[len(this.followers)-1].f = append(this.followers[len(this.followers)-1].f, followerId)
	this.m2[followeeId]++
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	for _, follower := range this.followers {
		if follower.id == followerId {
			for i, i2 := range follower.f {
				if i2 == followeeId {
					follower.f = append(follower.f[:i], follower.f[i+1:]...)
					break
				}
			}
			break
		}
	}
}
