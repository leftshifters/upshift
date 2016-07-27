package com.protocol10.taskapp;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

public class MainAdapter extends RecyclerView.Adapter<MainAdapter.MainHolder> {

    private List<String> list;
    private Context context;
    private OnItemClickListener onItemClickListener;

    public MainAdapter(Context context, List<String> list) {
        this.context = context;
        this.list = list;
    }

    @Override
    public MainHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        LayoutInflater layoutInflater = (LayoutInflater) context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        View view = layoutInflater.inflate(R.layout.row_main_content, parent, false);
        return new MainHolder(view);
    }

    @Override
    public void onBindViewHolder(MainHolder holder, int position) {
        holder.textView.setText(list.get(position));
    }

    @Override
    public int getItemCount() {
        return list.size();
    }

    public void setOnItemClickListener(OnItemClickListener onItemClickListener) {
        this.onItemClickListener = onItemClickListener;
    }

    public class MainHolder extends RecyclerView.ViewHolder {

        @BindView(R.id.text_view_title)
        TextView textView;

        public MainHolder(View itemView) {
            super(itemView);
            ButterKnife.bind(this, itemView);
            itemView.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {
                    if (onItemClickListener != null) {
                        onItemClickListener.onItemClick(getAdapterPosition(), v);
                    }
                }
            });
        }
    }
}
