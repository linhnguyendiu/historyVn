// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;
import "./LinkToken.sol";
import "./CertNFT.sol";

contract EduManage {
    address public owner;
    LINKToken public LINK;
    CertificateNFT public CertNFT;
    uint256 public postTokenReward = 1;
    uint256 public postPointToReward = 1000;
    uint256 public decimals = 18;

    struct Student {
        uint256 id;
        address stuAdd;
        string stu_name;
        uint256[] courses_id;
    }

    struct Grade {
        uint256 mark;
        bool isSet;
        bool isReceiveReward;
        bool isMintCert;
    }

    struct Post {
        uint256 id;
        address owner;
        uint256 point;
        uint256 rejectCount; 
        uint256 rewardCount;
    }

    struct Course {
        uint256 id;
        string name;
        uint256 users;
        uint256 price;
        uint256 reward;
        string courseType;
        string hashCourse;
        uint256 createTime;
    }

    uint256 studentCount;
    uint256 courseCount;
    uint256 postCount;

    mapping(uint256 => Student) students;
    mapping(uint256 => Course) courses;
    mapping(uint256 => Post) posts;
    mapping(address => bool) hasAccount;
    mapping(uint256 => mapping(uint256 => Grade)) public grades;

    constructor(address _tokenAddress, address _NFTAddress) {
        owner = msg.sender;
        LINK = LINKToken(_tokenAddress);
        CertNFT = CertificateNFT(_NFTAddress);
    }
    
    function addStudent(string memory name) public {
        require(!hasAccount[msg.sender], "Address already has account");
        studentCount++;
        students[studentCount].id = studentCount;
        students[studentCount].stu_name = name;
        students[studentCount].stuAdd = msg.sender;
        hasAccount[msg.sender] = true;
    }

    function addPost() public {
        require(hasAccount[msg.sender], "Not authorized to add post"); 
        postCount++;
        posts[postCount].id = postCount;
        posts[postCount].owner = msg.sender;
    }

    function addCourse(
        string memory name,
        uint256 price,
        uint256 reward,
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

    function checkEnrolledCourse(uint256 stu_id, uint256 course_id) public view returns (bool){
        require(hasAccount[msg.sender], "Not authorized to enroll course");
        bool isEnrolled = false;
        for (uint256 i = 0; i < students[stu_id].courses_id.length; i++) {
            if (students[stu_id].courses_id[i] == course_id) {
                isEnrolled = true;
                break;
            }
        }
        return isEnrolled;
    }

    function allowanceBuyCourse() public view returns (uint256) {
        return LINK.allowance(msg.sender, address(this));
    }

    function buyCourse(
        uint256 stu_id,
        uint256 course_id
        ) public payable {
        require(hasAccount[msg.sender] , "Not authorized to enroll course"); 
        require(!checkEnrolledCourse( stu_id, course_id), "You bought this course");
        require(courses[course_id].reward != 0,"Course does not exist");
            allowanceBuyCourse();
            LINK.transferFrom(msg.sender, address(this), courses[course_id].price*10**decimals);
            courses[course_id].users++;
            students[stu_id].courses_id.push(course_id);
    }

    function rewardToken(address recipient,uint256 _numberOfTokens) public payable{
        require(LINK.balanceOf(address(this)) >= _numberOfTokens*10**decimals, "Transaction fail");
        LINK.approve(address(this), _numberOfTokens*10**decimals);
        LINK.transferFrom(address(this), recipient, _numberOfTokens*10**decimals);
    }

    function submitGrade(
        uint256 stu_id,
        uint256 course_id,
        uint256 _mark 
        ) public returns (bool) {
            require(checkEnrolledCourse(stu_id, course_id), "Not authorized to submit exam");
            require(!grades[stu_id][course_id].isSet, "Your mark is update"); 
            grades[stu_id][course_id] = Grade(_mark, true, false, false);
        return true;
    }

    function checkAndTransferRewardCourse(
        uint256 stu_id,
        uint256 course_id,
        string memory image_uri_special,
        string memory image_uri_normal
        ) public payable {
            require(grades[stu_id][course_id].isSet, "Please take exam");
            require(!grades[stu_id][course_id].isReceiveReward, "You have received a reward for this course");
            require(!grades[stu_id][course_id].isMintCert, "You have received a certificate for this course");
            uint256 _mark = grades[stu_id][course_id].mark; 
            if (10 >=_mark && _mark >= 9) {
                rewardToken(msg.sender, courses[course_id].reward);
                CertNFT.mintCertificate(msg.sender, students[stu_id].stu_name, 
                courses[course_id].name, block.timestamp, courses[course_id].courseType, image_uri_special);
                grades[stu_id][course_id].isReceiveReward = true;
                grades[stu_id][course_id].isMintCert = true;
            } else if (_mark >= 7) {
                CertNFT.mintCertificate(msg.sender, students[stu_id].stu_name, 
                courses[course_id].name, block.timestamp, courses[course_id].courseType, image_uri_normal);
                grades[stu_id][course_id].isMintCert = true;
            } else {
                revert("Sorry, you have failed exam.");
            }
    }

    //tranferRewardPost  1000point => check + post - msg.sender, add post ?? address owner post + id post + point + check reject 
    function checkAndTransferRewardPost(
        uint256 post_id,
        uint256 point,
        uint256 reject_count
        ) public payable {
            posts[post_id].point = point;
            posts[post_id].rejectCount = reject_count;
            if (point == postPointToReward*(posts[post_id].rewardCount+1) && posts[post_id].rejectCount <= 3) {
                rewardToken(posts[post_id].owner, postTokenReward);
                posts[post_id].rewardCount++;
            }
            else {
                revert("Sorry, you don't have enough point to take reward token in this post or your post have a lot of reject");
            }
    }

    function getCourse(uint256 _id) public view returns (Course memory) {
        return (courses[_id]);
    }

    function getStudent(uint256 _id) public view returns (Student memory) {
        return (students[_id]);
    }

    function getCourseCount() public view returns (uint256) {
        return (courseCount);
    }

    function getGrades(uint256 _stu_id, uint256 _course_id) public view returns (Grade memory) {
        return (grades[_stu_id][_course_id]);
    }

    function getPosts(uint256 _id) public view returns (Post memory) {
        return (posts[_id]);
    }
}
