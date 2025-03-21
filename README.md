# Comp590 Assignment 6
# Reese Letts

Do you find your code showing race conditions? Try it with larger p values to get many concurrent processes. What can you do about that? Why or why not are you seeing race conditions?

Yes, when I ran Prog 2 with larger values of p, I encountered race conditions. The issue happens because multiple goroutines are trying to access and modify the same global variable, res, at the same time. This can lead to inconsistent results. To handle this, one approach is to use synchronization techniques like a mutex to ensure that only one goroutine can modify res at a time. This ensures that the changes to res happen in a controlled manner, preventing concurrent modifications.

Without a mutex, race conditions occur because the goroutines are attempting to update the global res variable concurrently. The lack of synchronization allows multiple goroutines to potentially read and write to res simultaneously, which can lead to incorrect results. By using a mutex, I was able to synchronize the goroutines, ensuring that only one can modify res at a time, which eliminates the race condition.
