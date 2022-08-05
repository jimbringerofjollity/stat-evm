// (c) 2022-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface Fit {
    function getFit(uint256[] memory v1) external view returns (uint256);
}

contract Test {
    uint256 public last;

    event Debug(string message, bytes32 res);

    Fit prec = Fit(0x0300000000000000000000000000000000000001);

    function testMe(uint256[] memory vals) public {
        last = prec.getFit(vals);
//        emit Debug("testMe()", last);
    }
}