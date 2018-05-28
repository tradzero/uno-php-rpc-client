<?php
// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// !
// Copyright 2018 acrazing <joking.young@gmail.com>
//
// @since 2018-05-24 09:06:59
// @version 1.0.0
// @desc messages.proto
namespace Uno;

/**
 */
class UnoClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Uno\PBEmpty $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Rent(\Uno\PBEmpty $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/uno.Uno/Rent',
        $argument,
        ['\Uno\UnoMessage', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Uno\UnoMessage $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Relet(\Uno\UnoMessage $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/uno.Uno/Relet',
        $argument,
        ['\Uno\PBEmpty', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Uno\UnoMessage $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Return(\Uno\UnoMessage $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/uno.Uno/Return',
        $argument,
        ['\Uno\PBEmpty', 'decode'],
        $metadata, $options);
    }

}
