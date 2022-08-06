// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface Median {
    function getMedian(uint256[] memory v1) external view returns (uint256);
}
interface Sampler {
    function sampleRandomNumber(
        uint256 distributionType,
        int256 a,
        int256 b,
        uint256 numSamples
    ) external view returns (int256[] calldata); // TODO: Should this be `memory`?
}
interface MatrixMult{
    function matrixMultiply(int256[][] memory a, int256[][] memory b) external view returns (int256[][] memory);
}

contract Test {
    uint256 public med;
    int256[] public sample;
    int256[][] product;
    event Debug(string message, int256 res);
    Median prec = Median(0x0300000000000000000000000000000000000001);
    Sampler sampler = Sampler(0x0300000000000000000000000000000000000004);
    MatrixMult matrixMult = MatrixMult(0x0300000000000000000000000000000000000005);

    function testMedian(uint256[] memory vals) public {
        med = prec.getMedian(vals);
    }
    
    function testSampler(uint256 distributionType, int256 a, int256 b, uint256 numSamples) public {
        sample = sampler.sampleRandomNumber(distributionType, a, b, numSamples);
    }

    function getLastSample() public view returns (int256[] memory) {
        return sample;
    }

    function testMatrixMult(int256[][] memory a, int256[][] memory b) public {
        product = matrixMult.matrixMultiply(a, b);
    }

    function getMatrixMulti() public view returns (int256[][] memory) {
        return product;
    }
}