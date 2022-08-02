// (c) 2022-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface Median {
    function getMedian() external view returns (bytes32);
}

contract Test {
    bytes32 public last;

    event Debug(string message, bytes32 res);

    Median prec = Median(0x0300000000000000000000000000000000000001);

    function testMe() public {
        //last = prec.getMedian();
        prec.getMedian();
//        emit Debug("testMe()", last);
    }
}