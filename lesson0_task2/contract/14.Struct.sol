// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

/*
    结构体（Structs）是一种用户自定义的数据类型，它允许你将多个变量组合成一个单一的类型
    格式：struct Name{
            uint id;
            string name;
            ...
        }
    示例：创建待办任务自定义类型，任务id和任务名称
             # 新建任务
             # 查询某一个任务
             # 修改任务名称
             # 完成任务     
*/
contract Struct {
    // 任务结构体
    struct TodoTask {
        string taskname;
        bool status;
    }

    // 函数修改器
    modifier requireIndex(uint256 _index) {
        require(taskList.length > 0 && _index < taskList.length, "index out of bound!");
        _;
    }

    // 机构体数组
    TodoTask[] public taskList;

    function create(string memory _name) external {
        taskList.push(TodoTask({taskname: _name, status: false}));
    }

    function getList() external view returns (TodoTask[] memory) {
        return taskList;
    }

    function updateTaskname(uint256 _index, string memory _name)
        external
        requireIndex(_index)
    {
        TodoTask storage task = taskList[_index];
        task.taskname = _name;
    }

    function completed(uint256 _index) external requireIndex(_index) {
        TodoTask storage task = taskList[_index];
        bool status = taskList[_index].status;
        if (!status) {
            taskList[_index].status = !taskList[_index].status;
        }
    }
}
