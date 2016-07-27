package com.protocol10.taskapp.creating;

/**
 * Created by akshay on 19/7/16.
 */

public interface Generic<T, R> {

    R onSuccess(T t);
}
