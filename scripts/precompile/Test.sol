// (c) 2022-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface Median {
    function getMedian(uint256 v1, uint256 v2, uint256 v3) external view returns (uint256);
}

contract Test {
    uint256 public last;

    event Debug(string message, bytes32 res);

    Median prec = Median(0x0300000000000000000000000000000000000001);

    function testMe(uint256 v1, uint256 v2, uint256 v3) public {
        last = prec.getMedian(v1, v2, v3);
//        emit Debug("testMe()", last);
    }
}