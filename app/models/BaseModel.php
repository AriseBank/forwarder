<?php 

use Illuminate\Validation\Validator;

class BaseModel extends Eloquent {


	protected static $rules = array();

	protected static $messages = array();

	protected $validator;

    public function __construct(array $attributes = array(), Validator $validator = null)
    {
        parent::__construct($attributes);

        $this->validator = $validator ?: \App::make('validator');
    }

    public function validate()
    {
        $v = $this->validator->make($this->attributes, static::$rules, static::$messages);

        if ($v->passes())
        {
            return $this;
        }

        $this->setErrors($v->messages());

        return $this;
    }

    protected function setErrors($errors)
    {
        $this->errors = $errors;
    }

    public function hasErrors()
    {
        return ! empty($this->errors);
    }

}
