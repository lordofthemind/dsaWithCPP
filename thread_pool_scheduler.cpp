#include <iostream>
#include <queue>
#include <vector>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <functional>
#include <atomic>

// Task structure with priority
struct Task
{
    int priority;
    std::function<void()> func;

    bool operator<(const Task &other) const
    {
        return priority < other.priority;
    }
};

class ThreadPool
{
public:
    ThreadPool(size_t num_threads);
    ~ThreadPool();
    void enqueueTask(int priority, std::function<void()> task);

private:
    std::vector<std::thread> workers;
    std::priority_queue<Task> tasks;
    std::mutex queue_mutex;
    std::condition_variable condition;
    std::atomic<bool> stop;
    void workerThread();
};

ThreadPool::ThreadPool(size_t num_threads) : stop(false)
{
    for (size_t i = 0; i < num_threads; ++i)
    {
        workers.emplace_back(&ThreadPool::workerThread, this);
    }
}

ThreadPool::~ThreadPool()
{
    stop = true;
    condition.notify_all();
    for (std::thread &worker : workers)
    {
        if (worker.joinable())
        {
            worker.join();
        }
    }
}

void ThreadPool::enqueueTask(int priority, std::function<void()> task)
{
    {
        std::lock_guard<std::mutex> lock(queue_mutex);
        tasks.push(Task{priority, task});
    }
    condition.notify_one();
}

void ThreadPool::workerThread()
{
    while (!stop)
    {
        Task task;
        {
            std::unique_lock<std::mutex> lock(queue_mutex);
            condition.wait(lock, [this]
                           { return stop || !tasks.empty(); });

            if (stop && tasks.empty())
                return;

            task = tasks.top();
            tasks.pop();
        }
        task.func();
    }
}

int main()
{
    ThreadPool pool(4);

    pool.enqueueTask(2, []()
                     { std::cout << "Task 1 (Priority 2)\n"; });
    pool.enqueueTask(1, []()
                     { std::cout << "Task 2 (Priority 1)\n"; });
    pool.enqueueTask(3, []()
                     { std::cout << "Task 3 (Priority 3)\n"; });
    pool.enqueueTask(5, []()
                     { std::cout << "Task 4 (Priority 5)\n"; });
    pool.enqueueTask(4, []()
                     { std::cout << "Task 5 (Priority 4)\n"; });

    std::this_thread::sleep_for(std::chrono::seconds(2));
    return 0;
}
