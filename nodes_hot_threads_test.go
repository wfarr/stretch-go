package stretch

import (
	"net/http"
	"testing"
)

func TestNodesHotThreads(t *testing.T) {
	response := `
::: [testcluster][someid][inet[/127.0.0.1:9300]]{data=true,master=false}

    0.1% (330micros out of 500ms) cpu usage by thread 'elasticsearch[testnode][scheduler][T#1]'
     10/10 snapshots sharing following 9 elements
       sun.misc.Unsafe.park(Native Method)
       java.util.concurrent.locks.LockSupport.parkNanos(LockSupport.java:226)
       java.util.concurrent.locks.AbstractQueuedSynchronizer$ConditionObject.awaitNanos(AbstractQueuedSynchronizer.java:2082)
       java.util.concurrent.ScheduledThreadPoolExecutor$DelayedWorkQueue.take(ScheduledThreadPoolExecutor.java:1090)
       java.util.concurrent.ScheduledThreadPoolExecutor$DelayedWorkQueue.take(ScheduledThreadPoolExecutor.java:807)
       java.util.concurrent.ThreadPoolExecutor.getTask(ThreadPoolExecutor.java:1068)
       java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1130)
       java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:615)
       java.lang.Thread.run(Thread.java:745)
`

	ts := testServer(http.StatusOK, response)
	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	hotThreads, err := cluster.GetHotThreads()

	if err != nil {
		t.Fail()
	}

	if hotThreads != response {
		t.Fail()
	}

}
