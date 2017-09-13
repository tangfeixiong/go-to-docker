<?php

namespace App\Repository;

use Illuminate\Support\Facades\DB;

use App\Battlefield;

class BattlefieldRepository
{
//
// ORM
//
    public function reapAll() {
        return \App\Battlefield::all();
    }  
       
    public function reapOne($id) {
        // return App\Battlefield::where("id", $id).get();
        return \App\Battlefield::find($id);
    }     
    
    /**
     * Create a new 'Battlefield' row.
     *
     * @param  Request  $request
     * @return Response
     */
    public function create($name, $desc, $matchItem, $status='not-ready', $payment='beta')
    {
        // Validate the ...

        $bf = new Battlefield;

        $bf->fill([
            'name' => $name,
            'description' => $desc,
            'match_item' => $matchItem,
            'status' => $status,
            'payment' => $payment
        ]);
        
        if ($bf->save())
            return $bf->id;
        
        // $bf = App\Battlefield::create([
        //     'name' => $name,
        //     'description' => $desc,
        //     'match_item' => $matchItem,
        //     'status' => $status,
        //     'payment' => $payment
        // ]);
        // if ($bf->save())
        //     return $bf->id;
    }
    
    public function toUpdate($id, $name, $desc, $matchItem, $status='not-ready', $payment='beta')
    {
        // Validate the ...

        $affectedRows = App\Battlefield::where('active', 1)
            ->update(['name' => $name,
                'description' => $desc,
                'match_item' => $matchItem,
                'status' => $status,
                'payment' => $payment
            ]);
            
        return $affectedRows;

        // $bf = new Battlefield;
        
        // $bf->where("id", $id);

        // $bf->name' = $name;
        // $bf->description = $desc;
        // $bf->match_item = $matchItem;
        // $bf->status = $status;
        // $bf->payment => $payment;
        
        // return $bf->save()
    }
    
    public function toDelete($id)
    {
        $bf = App\Battlefield::where('id', '==', $id);
        
        if ($bf.networkings().get().count() > 0)
            return NULL;
        if ($bf.checkalives().get().count() > 0)
            return NULL;
        
        $affectedRows = App\Battlefield::where('id', '==', $id)->delete();
        
        return $affectedRows;
    }

//
// Query Builder
//    
    /**
     * Show a list of all of 'Battlefield'.
     *
     * @return Response
     */
    public function index()
    {
        $rows = DB::table('battlefield')->get();

        // return view('battlefield.index', ['battlefield' => $bf]);
        return $rows;
    }
    
    public function selectOne($id)
    {
        $row = DB::table('battlefield')->where('id', '==', $id)->get();

        return $row;
    }
    
    public function insert($name, $desc, $matchItem, $status='not-ready', $payment='beta')
    {
        $id = DB::table('battlefield')
            ->insert([
                [
                    'name' => $name,
                    'description' => $desc,
                    'match_item' => $matchItem,
                    'status' => $status,
                    'payment' => $payment
                ],
            ]);        
        return $id;
    }
    
    public function update($id, $name, $desc, $matchItem, $status='not-ready', $payment='beta')
    {
        $affectedRows = DB::table('battlefield')
            ->where('id', $id)
            ->insert([
                [
                    'name' => $name,
                    'description' => $desc,
                    'match_item' => $matchItem,
                    'status' => $status,
                    'payment' => $payment
                ],
            ]);        
        return $affectedRows;
    }
    
    public function delete($id)
    {
        $rows = DB::table('networking')->select('battlefield_id')->where('battlefield_id', '==', $id)->get();
        if ($rows.count() > 0)
            return NULL;

        $rows = DB::table('check_alive')->select('battlefield_id')->where('battlefield_id', '==', $id)->get();
        if ($rows.count() > 0)
            return NULL;
        
        $affectedRows = DB::table('battlefield')->where('id', '==', $id)->delete();
        
        return $affectedRows;
    }
}