// tranferRewardCourse, mintCertificate, buyCourse, tranferRewardPost  
//struct course : id, hash_course, reward, fee, name
// struct student : id, name, list_course, add, 
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "./LINKToken.sol";

contract eduManage {
    address public owner;
    LINKToken public LINK;
    struct Student {
        int256 id;
        int256[] courses_id;
        int256[] marks;
    }

    struct Course {
        int256 id;
        string name;
        int256 users;
        uint256 price;
        int256 reward;
        string courseType;
        string hashCourse;
        uint256 createTime;
    }

    int256 studentCount;
    int256 courseCount;

    mapping(int256 => Student) students;
    mapping(int256 => Course) courses;
    mapping(address => bool) hasAccount;

    constructor(address _tokenAddress) {
        owner = msg.sender;
        LINK = LINKToken(_tokenAddress);
    }

    function addStudent() public {
        require(!hasAccount[msg.sender], "Address already has account");
        studentCount++;
        students[studentCount].id = studentCount;
        hasAccount[msg.sender] = true;
    }

    function addCourse(
        string memory name,
        uint256 price,
        int256 reward,
        string memory courseType,
        string memory hashCourse
    ) public {
        require(msg.sender == owner, "Not authorized to add course"); 
        courseCount++;
        courses[courseCount].id = courseCount;
        courses[courseCount].name = name;
        courses[courseCount].price = price;
        courses[courseCount].reward = reward;
        courses[courseCount].courseType = courseType;
        courses[courseCount].hashCourse = hashCourse;
        courses[courseCount].createTime = block.timestamp;
    }

    function buyCourse(
        int256 stu_id,
        int256 course_id
        ) public {
        require(!hasAccount[msg.sender] , "Not authorized to buy course"); 
        courses[stu_id].users++;
        students[stu_id].courses_id.push(course_id);
        LINK.transfer(address(this), courses[course_id].price);
    }

    function rewardToken(address recipient,uint256 _numberOfTokens) public payable{
        require(LINK.balanceOf(address(this)) >= _numberOfTokens);
        LINK.approve(address(this), _numberOfTokens);
        LINK.transferFrom(address(this), recipient, _numberOfTokens);
    }

    function getCourse(int256 _id) public view returns (Course memory) {
        return (courses[_id]);
    }

    function getStudent(int256 _id) public view returns (Student memory) {
        return (students[_id]);
    }


    function getCourseCount() public view returns (int256) {
        return (courseCount);
    }
}
