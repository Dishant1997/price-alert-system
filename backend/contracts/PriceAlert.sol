// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PriceAlert {
    struct Alert {
        address user;
        string asset; // e.g., BTC/USD
        uint256 priceThreshold;
        bool above; // true if alert is triggered when price is above the threshold
    }

    Alert[] public alerts;
    event NewAlert(address indexed user, string asset, uint256 priceThreshold, bool above);

    function createAlert(string memory _asset, uint256 _priceThreshold, bool _above) public {
        alerts.push(Alert(msg.sender, _asset, _priceThreshold, _above));
        emit NewAlert(msg.sender, _asset, _priceThreshold, _above);
    }

    function getAlerts() public view returns (Alert[] memory) {
        return alerts;
    }
}