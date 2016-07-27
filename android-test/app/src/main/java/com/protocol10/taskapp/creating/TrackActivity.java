package com.protocol10.taskapp.creating;

import android.database.Cursor;
import android.os.Bundle;
import android.provider.MediaStore;
import android.support.annotation.Nullable;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.util.Log;

import com.protocol10.taskapp.R;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeUnit;

import butterknife.BindView;
import butterknife.ButterKnife;
import rx.Observable;
import rx.Observer;
import rx.Subscriber;
import rx.Subscription;
import rx.android.schedulers.AndroidSchedulers;
import rx.functions.Func0;
import rx.schedulers.Schedulers;
import rx.subjects.AsyncSubject;
import rx.subjects.PublishSubject;
import rx.subjects.ReplaySubject;
import rx.subjects.Subject;
import rx.subscriptions.CompositeSubscription;

public class TrackActivity extends AppCompatActivity {

    @BindView(R.id.recycler_view)
    RecyclerView recyclerView;

    private LinearLayoutManager linearLayoutManager;

    private List<TrackModel> list;

    private TrackAdapters adapters;
    private CompositeSubscription compositeSubscription;
    /**
     * PROJECTION - Retrieves the value for Tracks from MediaStore DataBase in
     * android Only used for MediaProjection for general tracks
     */

    private static final String[] TRACKS_COLUMNS = {
            MediaStore.Audio.Media._ID, MediaStore.Audio.Media.DATA,
            MediaStore.Audio.Media.TITLE, MediaStore.Audio.Media.ARTIST,
            MediaStore.Audio.Media.ALBUM, MediaStore.Audio.Media.DURATION,
            MediaStore.Audio.Media.ALBUM_ID, MediaStore.Audio.Media.ARTIST_ID};

    private Subscription subscriptions;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_tracks);
        ButterKnife.bind(this);
        list = new ArrayList<>();
        linearLayoutManager = new LinearLayoutManager(getApplicationContext());
        recyclerView.setLayoutManager(linearLayoutManager);
        adapters = new TrackAdapters(list, getApplicationContext());
        recyclerView.setAdapter(adapters);
        Observable<Long> longObservable = Observable.interval(1000, TimeUnit.MILLISECONDS);
        Subject<Long, Long> longSubject = PublishSubject.create();
        longObservable.subscribe(longSubject);

        longSubject.subscribe(new Observer<Long>() {
            @Override
            public void onCompleted() {

            }

            @Override
            public void onError(Throwable e) {

            }

            @Override
            public void onNext(Long aLong) {
                Log.i("Observable", "Observable" + aLong);
            }
        });
        longSubject.subscribe(new Observer<Long>() {
            @Override
            public void onCompleted() {

            }

            @Override
            public void onError(Throwable e) {

            }

            @Override
            public void onNext(Long aLong) {
                Log.i("Observable2", "Observable" + aLong);
            }
        });
        longSubject.onNext(50l);
        adapters.notifyDataSetChanged();
        updateRecyclerView();
//        compositeSubscription = new CompositeSubscription();
//        ReplaySubject<String> subject = ReplaySubject.create();
//        subject.onNext("James");
//        subject.onNext("Sir Arthur");
//        subject.onNext("Jamie");
//        Subscription subscription = subject.subscribe(getDefaultSubscriber());
//        compositeSubscription.add(subscription);
//        subject.onNext("Tyrion");
//        subject.onCompleted();
//        Subscription subscription1 = subject.subscribe(getTardySubscriber());
//        compositeSubscription.add(subscription1);
        compositeSubscription = new CompositeSubscription();
        ReplaySubject<String> stockReplaySubject = ReplaySubject.create();

        stockReplaySubject.onNext("A");
        stockReplaySubject.onNext("B");
        stockReplaySubject.onNext("C");

        Subscription defaultSub = stockReplaySubject.subscribe(getDefaultSubscriber());
        compositeSubscription.add(defaultSub); // All three values will be delivered.

        stockReplaySubject.onNext("D");
        stockReplaySubject.onCompleted(); // Terminate the stream with a completed event.

        // Subscribe again, this time the subscriber will get all events and the terminal event
        // right away. All items are "replayed" even though the stream has completed.
        Subscription tardySubscription = stockReplaySubject.subscribe(getTardySubscriber());
        compositeSubscription.add(tardySubscription);

    }

    private Subscriber<String> getDefaultSubscriber() {
        return new Subscriber<String>() {
            @Override
            public void onCompleted() {
                Log.i("GET_DEFAULT_SUBSCRIBER", "onCompleted");
            }

            @Override
            public void onError(Throwable e) {
                Log.i("GET_DEFAULT_SUBSCRIBER", "onError " + e.getLocalizedMessage());
            }

            @Override
            public void onNext(String s) {
                Log.i("GET_DEFAULT_SUBSCRIBER", "onNext " + s);
            }
        };
    }

    private Subscriber<String> getTardySubscriber() {
        return new Subscriber<String>() {
            @Override
            public void onCompleted() {
                Log.i("getTardySubscriber", "onCompleted");
            }

            @Override
            public void onError(Throwable e) {
                Log.i("getTardySubscriber", "onError " + e.getLocalizedMessage());
            }

            @Override
            public void onNext(String s) {
                Log.i("getTardySubscriber", "onNext " + s);
            }
        };
    }

    /**
     * A long running operation that fetch all the AudioTracks using @{@link MediaStore } framework
     * Consider using it in a actual device rather than emulator
     *
     * @return
     */
    public List<TrackModel> retrieveTracks() throws IOException {
        List<TrackModel> trackModelList = new ArrayList<>();
        Cursor cursor = getContentResolver().query(MediaStore.Audio.Media.EXTERNAL_CONTENT_URI,
                TRACKS_COLUMNS, "", null, MediaStore.Audio.Media.TITLE + " asc");
        if (cursor == null) {
            throw new IOException("Unable to fetch Media");
        } else if (!cursor.moveToFirst()) {
            throw new IOException("No Media found on device");
        } else {
            do {
                String id = cursor.getString(cursor.getColumnIndex(MediaStore.Audio.Media._ID)); // ID
                String path = cursor.getString(cursor
                        .getColumnIndex(MediaStore.Audio.Media.DATA));// DATA
                String title = cursor.getString(cursor
                        .getColumnIndex(MediaStore.Audio.Media.TITLE));// TITLE

                trackModelList.add(new TrackModel(id, title, path));
            } while (cursor.moveToNext());

        }
        return trackModelList;
    }


    Observable<List<TrackModel>> updateRecyclerView() {
        /**
         * This operator @{@link Observable.defer} is great operator, this helps to avoid any execution
         * of task unless any subsciber is not attached this will not react or produce any result
         */
        return Observable.defer(new Func0<Observable<List<TrackModel>>>() {
            @Override
            public Observable<List<TrackModel>> call() {
                try {
                    return Observable.just(retrieveTracks());// create an observable that retruns list
                } catch (IOException e) {
                    return Observable.error(e);// Call onError when any Exception is thrown, this will help to call onError in Subscriber class
                }
            }
        });
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        if (subscriptions != null && subscriptions.isUnsubscribed()) {
            subscriptions.isUnsubscribed();
        }
    }
}
