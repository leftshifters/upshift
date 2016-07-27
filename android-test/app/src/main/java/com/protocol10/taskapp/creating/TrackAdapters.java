package com.protocol10.taskapp.creating;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import com.protocol10.taskapp.R;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

public class TrackAdapters extends RecyclerView.Adapter<TrackAdapters.TracksHolder> {

    private List<TrackModel> list;
    private Context context;

    public TrackAdapters(List<TrackModel> list, Context context) {
        this.list = list;
        this.context = context;
    }

    @Override
    public TracksHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        LayoutInflater layoutInflater = (LayoutInflater) context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        View view = layoutInflater.inflate(R.layout.row_tracks, parent, false);
        return new TracksHolder(view);
    }

    @Override
    public void onBindViewHolder(TracksHolder holder, int position) {
        TrackModel trackModel = list.get(position);
        holder.txtTrackId.setText(String.valueOf(trackModel.id));
        holder.txtTrackId.setText(trackModel.trackName);
        Log.i(TrackAdapters.class.getName(), trackModel.toString());
    }

    @Override
    public int getItemCount() {
        return list.size();
    }

    public class TracksHolder extends RecyclerView.ViewHolder {

        @BindView(R.id.text_track_id)
        TextView txtTrackId;
        @BindView(R.id.text_track_name)
        TextView txtTrackName;

        public TracksHolder(View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
        }
    }
}
