// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "./LINKToken.sol";
import "./CertNFT.sol";

contract EduManage {
    address public owner;
    LINKToken public LINK;
    CertificateNFT public CertNFT;
    uint256 public postTokenReward = 1;
    uint256 public postPointToReward = 10;
    uint256 public decimals = 8;

    struct Student {
        uint256 id;
        address stuAdd;
        string stu_name;
        uint256[] courses_id;
    }

    struct Grade {
        uint256 mark;
        string hashResultExam;
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

    event CertificateMinted(
        address recipient,
        uint256 tokenId,
        string name,
        string course,
        uint256 date,
        string cerType,
        string imageUri
    );

    constructor(address _tokenAddress, address _NFTAddress) {
        owner = msg.sender;
        LINK = LINKToken(_tokenAddress);
        CertNFT = CertificateNFT(_NFTAddress);
    }
    
    function addStudent(address stu_add, uint256 stu_id, string memory name) public {
        require(!hasAccount[stu_add], "Address already has account");
        studentCount++;
        students[stu_id].id = stu_id;
        students[stu_id].stu_name = name;
        students[stu_id].stuAdd = stu_add;
        hasAccount[stu_add] = true;
    }

    function addPost(uint256 post_id, address owner_add) public {
        require(hasAccount[owner_add], "Not authorized to add post"); 
        require(posts[post_id].id == 0, "Post had created"); 
        postCount++;
        posts[post_id].id = post_id;
        posts[post_id].owner = owner_add;
    }

    function addCourse(
        uint256 course_id,
        string memory name,
        uint256 price,
        uint256 reward,
        string memory courseType,
        string memory hashCourse
    ) public {
        require(msg.sender == owner, "Not authorized to add course"); 
        require(courses[course_id].id == 0, "Course had created"); 
        courseCount++;
        courses[course_id].id = course_id;
        courses[course_id].name = name;
        courses[course_id].price = price;
        courses[course_id].reward = reward;
        courses[course_id].courseType = courseType;
        courses[course_id].hashCourse = hashCourse;
        courses[course_id].createTime = block.timestamp;
    }

    function checkEnrolledCourse(uint256 stu_id, uint256 course_id) public view returns (bool){
        require(hasAccount[students[stu_id].stuAdd], "Not authorized to enroll course");
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
        require(hasAccount[students[stu_id].stuAdd] , "Not authorized to enroll course"); 
        require(!checkEnrolledCourse(stu_id, course_id), "You bought this course");
        require(courses[course_id].reward != 0,"Course does not exist");
            allowanceBuyCourse();
            LINK.transferFrom(students[stu_id].stuAdd, address(this), courses[course_id].price*10**decimals);
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
        uint256 _mark ,
        string memory hash_result_exam
        ) public returns (bool) {
            require(checkEnrolledCourse(stu_id, course_id), "Not authorized to submit exam");
            require(!grades[stu_id][course_id].isSet, "Your mark is update"); 
            grades[stu_id][course_id] = Grade(_mark, hash_result_exam, true, false, false);
        return true;
    }

    function checkAndTransferRewardCourse(
        uint256 stu_id,
        uint256 course_id,
        uint256 token_id,
        string memory image_uri
        ) public payable {
            require(grades[stu_id][course_id].isSet, "Please take exam");
            require(!grades[stu_id][course_id].isReceiveReward, "You have received a reward for this course");
            require(!grades[stu_id][course_id].isMintCert, "You have received a certificate for this course");
            uint256 _mark = grades[stu_id][course_id].mark; 
            if (10 >=_mark && _mark >= 9) {
                rewardToken(students[stu_id].stuAdd, courses[course_id].reward);
                CertNFT.mintCertificate(token_id, students[stu_id].stuAdd, students[stu_id].stu_name, 
                courses[course_id].name, block.timestamp, courses[course_id].courseType, image_uri);
                emit CertificateMinted(students[stu_id].stuAdd, token_id, students[stu_id].stu_name, 
                courses[course_id].name, block.timestamp, courses[course_id].courseType, image_uri);
                grades[stu_id][course_id].isReceiveReward = true;
                grades[stu_id][course_id].isMintCert = true;
            } else if (_mark >= 7) {
                CertNFT.mintCertificate(token_id, students[stu_id].stuAdd, students[stu_id].stu_name, 
                courses[course_id].name, block.timestamp, courses[course_id].courseType, image_uri);
                emit CertificateMinted(students[stu_id].stuAdd, token_id, students[stu_id].stu_name, 
                courses[course_id].name, block.timestamp, courses[course_id].courseType, image_uri);
                grades[stu_id][course_id].isMintCert = true;
            } else {
                revert("Sorry, you have failed exam.");
            }

    }

    //tranferRewardPost  1000point => check + post - msg.sender, add post ?? address owner post + id post + point + check reject 
    function checkAndTransferRewardPost(
        uint256 post_id,
        uint256 point
        ) public payable {
            if ( posts[post_id].rewardCount == 0) {
                if (point/ postPointToReward >= 1 ) {
                    posts[post_id].point = point;
                    rewardToken(posts[post_id].owner, postTokenReward*(point/postPointToReward));
                    posts[post_id].rewardCount++;
                }
                else {
                revert("Sorry, you don't have enough point to take reward token in this post");
            }
            } 
            else {
                if ((point > posts[post_id].point)&&((point-posts[post_id].point)/ postPointToReward >= 1 ) ){
                    rewardToken(posts[post_id].owner, postTokenReward*(point-posts[post_id].point)/postPointToReward);
                    posts[post_id].point = point;
                    posts[post_id].rewardCount++;
                }
                else {
                revert("Sorry, you don't have enough point to take reward token in this post");
            }
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
