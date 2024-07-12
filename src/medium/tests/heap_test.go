package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestKClosest(t *testing.T) {
    assert.ElementsMatch(t, [][]int{{3,3},{-2,4}}, m.KClosest([][]int{{3,3}, {5,-1}, {-2,4}}, 2))
    assert.ElementsMatch(t, [][]int{{-5,4},{4,6}}, m.KClosest([][]int{{-5,4}, {-6,-5}, {4,6}}, 2))
    assert.ElementsMatch(t, 
        [][]int{{0,1},{-2,0},{-3,2},{0,-5},{-5,4},{-4,-6}}, 
        m.KClosest([][]int{{-5,4},{-3,2},{0,1},{-3,7},{-2,0},{-4,-6},{0,-5}}, 6),
    )
}

func TestFindKthLargest(t *testing.T) {
    assert.Equal(t, 3, m.FindKthLargest([]int{7,6,5,4,3,2,1}, 5))
}

func TestLeastInterval(t *testing.T) {
    assert.Equal(t, 8, m.LeastInterval([]byte{'A','A','A','B','B','B'}, 2))
    assert.Equal(t, 6, m.LeastInterval([]byte{'A','C','A','B','D','B'}, 1))
}

func TestTwitter(t *testing.T) {
    twtr := m.TwitterConstructor()
    twtr.PostTweet(1, 5) // User 1 posts a new tweet (id = 5).
    assert.Equal(t, []int{5}, twtr.GetNewsFeed(1))  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
    twtr.Follow(1, 2)    // User 1 follows user 2.
    twtr.PostTweet(2, 6) // User 2 posts a new tweet (id = 6).
    assert.Equal(t, []int{6, 5}, twtr.GetNewsFeed(1))  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
    twtr.Unfollow(1, 2)  // User 1 unfollows user 2.
    assert.Equal(t, []int{5}, twtr.GetNewsFeed(1))  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.

    twtr1 := m.TwitterConstructor()
    twtr1.Follow(1, 5)
    assert.Equal(t, []int{}, twtr1.GetNewsFeed(1))

    twtr2 := m.TwitterConstructor()
    twtr2.PostTweet(1,5)
    twtr2.Follow(1,2)
    twtr2.Follow(2,1)
    assert.Equal(t, []int{5}, twtr2.GetNewsFeed(2))
    twtr2.PostTweet(2,6)
    assert.Equal(t, []int{6,5}, twtr2.GetNewsFeed(1))
    assert.Equal(t, []int{6,5}, twtr2.GetNewsFeed(2))
    twtr2.Unfollow(2,1)
    assert.Equal(t, []int{6,5}, twtr2.GetNewsFeed(1))
    assert.Equal(t, []int{6}, twtr2.GetNewsFeed(2))
    twtr2.Unfollow(1,2)
    assert.Equal(t, []int{5}, twtr2.GetNewsFeed(1))
    assert.Equal(t, []int{6}, twtr2.GetNewsFeed(2))
}

