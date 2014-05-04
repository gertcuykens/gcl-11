package my.endpoints;

import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.*;
import android.view.View.OnClickListener;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

public class MainActivity extends Activity {
    public static Activity activity;
    EditText graphValue;
    EditText messageValue;
    Button submitButton;
    Button publishButton;
    TextView textViewResults;

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.gcl11);
        activity = this;
        
        this.messageValue = (EditText) findViewById(R.id.messageValue);
        this.graphValue = (EditText) findViewById(R.id.graphValue);

        this.submitButton = (Button) findViewById(R.id.submitButton);
        this.submitButton.setOnClickListener(new OnClickListener() {
            public void onClick(View view) {onClickSubmit(view);}
        });

        this.publishButton = (Button) findViewById(R.id.publishButton);
        this.publishButton.setOnClickListener(new OnClickListener() {
            public void onClick(View view) {onClickRequest(view);}
        });

        this.textViewResults = (TextView) findViewById(R.id.textViewResults);
        this.textViewResults.setOnClickListener(new OnClickListener() {
            public void onClick(View view) {
                onClickList(view);
            }
        });

        Intent intent = new Intent(this, LoginUsingLoginFragmentActivity.class);
        startActivity(intent);
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        MenuInflater inflater = getMenuInflater();
        inflater.inflate(R.menu.login, menu);
        return super.onCreateOptionsMenu(menu);
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        Intent intent = new Intent(this, LoginUsingLoginFragmentActivity.class);
        startActivity(intent);
        return true;
    }

    private void onClickSubmit(View view) {
        Global g = Global.getInstance();
        g.setMessage(messageValue.getText().toString());
        Context context = view.getContext();
        new SubmitTask().execute(context);
    }

    private void onClickList(View view) {
        Context context = view.getContext();
        new ListTask().execute(context);
    }

    private void onClickRequest(View view) {
        Global g = Global.getInstance();
        g.setMessage(messageValue.getText().toString());
        g.setGraph(graphValue.getText().toString());
        Context context = view.getContext();
        new RequestTask().execute(context);
    }

    @Override
    public void onActivityResult(int requestCode, int resultCode, Intent data) {
        Log.i("graph", "----------------------------");
        Log.i("graph", String.valueOf("MainActivity onActivityResult:"+resultCode));
        Log.i("graph", "----------------------------");
    }

    static void toaster(Context c, String s) {
        Toast toast = Toast.makeText(c, s, Toast.LENGTH_SHORT);
        toast.setGravity(Gravity.TOP|Gravity.RIGHT, 0, 0);
        toast.show();
    }

}

/*
    @Override
    public void onActivityResult(int requestCode, int resultCode, Intent data) {
        new ListTask().execute(this);
    }

    static void alert(Context c, String s) {
        new AlertDialog.Builder(c)
                .setTitle(R.string.login_failed_dialog_title)
                .setMessage(s)
                .setPositiveButton(R.string.ok_button, null)
                .show();
    }
*/

//import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
//import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;

/*
switch (item.getItemId()) {
    case R.id.action_settings:
        Intent intent = new Intent(this, LoginUsingLoginFragmentActivity.class);
        startActivity(intent);
        return true;
    default:
        return super.onOptionsItemSelected(item);
}
*/
