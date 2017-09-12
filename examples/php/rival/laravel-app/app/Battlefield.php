<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Battlefield extends Model
{
    /**
     * The connection name for the model.
     *
     * @var string
     */
    protected $connection = 'mysql_dev';

    /**
     * The table associated with the model.
     *
     * @var string
     */
    protected $table = 'battlefield';


    /**
     * The primary key for the model.
     *
     * @var string
     */
    // protected $primaryKey = 'id';
    

    /**
     * Indicates if the IDs are auto-incrementing.
     *
     * @var bool
     */
    // public $incrementing = true;

    /**
     * Indicates if the model should be timestamped.
     *
     * @var bool
     */
    public $timestamps = false;

    /**
     * The storage format of the model's date columns.
     *
     * @var string
     */
    protected $dateFormat = 'U';

    const CREATED_AT = 'created_at';
    const UPDATED_AT = 'last_updated';

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = ['name', 'description', 'match_item', 'status', 'payment'];

}
