pragma solidity ^0.4.19;

import "./Pool.sol";

contract Pool {
  string publicKey;

  //these will be encrypted in the future and are STC
  string name;
  string email;
  string bio;
  string ip_address;
  string location;
  address wallet_address;
  int status; // 0 = rejected, 1 = approved, 2 = pending //in the future this should be a mapping(address=>int) for multiple pools
  bool exists;
  mapping pools (address=>string); //pools to public keys
  address owner;
}

function Node(
  string _publicKey,
  string _name,
  string _email,
  string _bio,
  string _ip_address,
  string _location){

    publicKey  = _publicKey;
    name       = _name;
    email      = _email;
    bio        = _bio;
    ip_address = _ip_address;
    location   = _location;
    status     = 2;
    exists     = 1;
    owner      = msg.sender;

}

function getPublicKey() public returns (string){
  return publicKey;
}

function getName() public returns (string){
  return name;
}

function getEmail() public returns (string){
  return email;
}

function getBio() public returns (string){
  return bio;
}

function getIP() public returns (string){
  return ip_address;
}

function getLocation() public returns (string){
  return location;
}

function doesExist() public returns (bool){
  return exists;
}

//setters

function setPublicKey(string _publicKey) public {
  require(msg.sender == owner);
  publicKey = _publicKey;
}

function setName(string _name) public {
  require(msg.sender == owner);
  name = _name;
}

function setEmail(string _email) public {
  require(msg.sender == owner);
  email = _email;
}

function setBio(string _bio) public {
  require(msg.sender == owner);
  bio = _bio;
}

function setIP(string _ip_address) public {
  require(msg.sender == owner);
  ip_address = _ip_address;
}

function setLocation(string _location) public {
  require(msg.sender == owner);
  location = _location;
}

function addPool() external public{
  Pool p = Pool.at(msg.sender);
  require(p.getNode(address(this)).doesExist());
  pools[msg.sender] = p.getPublicKey();
}
