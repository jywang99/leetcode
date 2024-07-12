package medium

import (
	"math"

	com "jy.org/leetcode/src/common"
)

// 973. K Closest Points to Origin
func KClosest(points [][]int, k int) [][]int {
    getDst := func(p []int) float64 {
        fx := float64(p[0])
        fy := float64(p[1])
        return float64(math.Sqrt(math.Pow(fx, 2) + math.Pow(fy, 2)))
    }

    hp := com.NewHeap(points, func(pa, pb []int) int {
        ad, bd := getDst(pa), getDst(pb)
        if ad < bd {
            return -1
        }
        return 1
    })

    res := make([][]int, 0)
    for i:=0; i<k; i++ {
        res = append(res, hp.PopTop())
    }

    return res
}

// 215. Kth Largest Element in an Array
func FindKthLargest(nums []int, k int) int {
    hp := com.NewHeap(nums, func(a, b int) int {
        if a == b {
            return 0
        }
        if a > b {
            return 1
        }
        return -1
    })
    for hp.GetSize() > k {
        hp.PopTop()
    }
    return hp.GetTop()
}

// 621. Task Scheduler
func LeastInterval(tasks []byte, n int) int {
    type tnode struct {
        char byte
        freq int
        next int
    }

    // frequency map
    tmap := make(map[byte]int)
    for _, t := range tasks {
        v, e := tmap[t]
        if !e {
            tmap[t] = 1
            continue
        }
        tmap[t] = v + 1
    }

    // heap, frequent nodes prioritized
    thp := com.NewHeap([]*tnode{}, func(a, b *tnode) int {
        if a.freq == b.freq {
            return 0
        }
        if a.freq > b.freq {
            return 1
        }
        return -1
    })
    for t, f := range tmap {
        thp.Insert(&tnode{
            char: t,
            freq: f,
        })
    }
    tq := []*tnode{}

    // start clock
    time := 0
    for thp.GetSize() > 0 || len(tq) > 0 {
        // check if any task in queue is ready again
        if len(tq) > 0 {
            qhead := tq[0]
            if qhead.next == time {
                tq = tq[1:]
                thp.Insert(qhead)
            }
        }

        // only if there's task to process at this time
        if thp.GetSize() > 0 {
            // process task
            task := thp.PopTop()
            task.freq --

            // if not the last occurrence, push it to queue
            task.next = time + n + 1
            if task.freq > 0 {
                tq = append(tq, task)
            }
        }

        time ++
    }

    return time
}

// 355. Design Twitter
type tweet struct {
    id int
    tstamp int
}

type Twitter struct {
    userPosts map[int][]tweet
    userFollows map[int]map[int]bool
    numTweets int
    time int
}

func TwitterConstructor() Twitter {
    return Twitter{
        userPosts: make(map[int][]tweet),
        userFollows: make(map[int]map[int]bool),
        numTweets: 10,
        time: 0,
    }
}

// user always follows himself
func (this *Twitter) followSelf(userId int)  {
    if this.userFollows[userId] == nil {
        ff := make(map[int]bool)
        this.userFollows[userId] = ff
    }
    this.userFollows[userId][userId] = true
}

func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.followSelf(userId)

    tweets, e := this.userPosts[userId]
    if !e {
        tweets = make([]tweet, 0)
    }
    tweets = append(tweets, tweet{
        id: tweetId,
        tstamp: this.time,
    })
    this.userPosts[userId] = tweets
    this.time ++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
    type post struct {
        userId int
        postId int
        idx int
        tstamp int
    }

    // get one latest post for each user this user follows, and himself
    ffs := this.userFollows[userId]
    latest := make([]post, 0)
    for fuid, fol := range ffs {
        if !fol {
            continue
        }
        posts := this.userPosts[fuid]
        lidx := len(posts)-1
        if lidx < 0 {
            continue
        }
        p := posts[lidx]
        latest = append(latest, post{
            userId: fuid,
            postId: p.id,
            tstamp: p.tstamp,
            idx: lidx,
        })
    }

    hp := com.NewHeap(latest, func(a, b post) int {
        if a.tstamp == b.tstamp {
            return 0
        }
        if a.tstamp < b.tstamp {
            return 1
        }
        return -1
    })

    res := make([]int, 0, this.numTweets)
    
    for i:=0; i<this.numTweets && hp.GetSize()>0; i++ {
        // latest = heap top
        p := hp.PopTop()
        res = append(res, p.postId)

        // add next latest post from the same user
        nidx := p.idx - 1
        // that was the last one
        if nidx < 0 {
            continue
        }
        np := this.userPosts[p.userId][nidx]
        p = post{
            userId: p.userId,
            postId: np.id,
            tstamp: np.tstamp,
            idx: nidx,
        }
        hp.Insert(p)
    }

    return res
}

func (this *Twitter) Follow(followerId int, followeeId int)  {
    this.followSelf(followerId)

    ffs, e := this.userFollows[followerId]
    if !e {
        ffs = make(map[int]bool)
    }
    ffs[followeeId] = true
    this.userFollows[followerId] = ffs
}

func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if _, e := this.userFollows[followerId]; !e {
        return
    }
    this.userFollows[followerId][followeeId] = false
}

