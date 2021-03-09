<?php namespace github_com\xiclonn\php\etp\err;

class Error {
// section 1 starts:	
	static function Nil () { return new self (null); }

	function __construct ($descrp, $sec = null) {
		$this->descrp = $descrp;
		$this->sec = $sec;
	}

// section 2 starts:
	private $descrp;
	private $sec;

// section 3 starts:
	function Descrp () {
		if ($this->descrp == null) {
			return "";
		}
		return $this->descrp;
	}
	function Sec () {
		if ($this->sec == null) {
			return self::Nil ();
		}
		return $this->sec;
	}

}
