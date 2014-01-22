package com.appspot;

import android.accounts.AccountManager;
import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.util.Pair;
import android.view.Gravity;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

//import com.google.android.gms.common.AccountPicker;
//import com.google.android.gms.common.GooglePlayServicesUtil;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;

import java.util.Arrays;

public class MainActivity extends Activity implements View.OnClickListener {
    //private GoogleAccountCredential user;
    private GoogleAccountCredential user2;
    private TextView userStatus;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.endpoints);

        //String AUDIENCE = "server:client_id:1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com";
        // user = GoogleAccountCredential.usingAudience(this, AUDIENCE);

        String SCOPE="https://www.googleapis.com/auth/devstorage.read_only";
        user2=GoogleAccountCredential.usingOAuth2(this, Arrays.asList(SCOPE.split(" ")));

        userStatus = (TextView) findViewById(R.id.userStatus);

        Button userButton = (Button) findViewById(R.id.userButton);
        userButton.setOnClickListener(this);

        Button storageButton = (Button) findViewById(R.id.storageButton);
        storageButton.setOnClickListener(this);

        Button serviceButton = (Button) findViewById(R.id.serviceButton);
        serviceButton.setOnClickListener(this);

        Button productButton = (Button) findViewById(R.id.productButton);
        productButton.setOnClickListener(this);
    }

    @Override
    public void onClick(View view) {
        Context context = view.getContext();
        switch(view.getId()) {
            case R.id.userButton: startActivityForResult(user2.newChooseAccountIntent(), 1); break;
            case R.id.storageButton: new StorageTask().execute(Pair.create(context, user2)); break;
            case R.id.serviceButton: new ServiceTask().execute(Pair.create(context, user2)); break;
            case R.id.productButton: new ProductTask().execute(Pair.create(context, user2)); break;
        }
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        switch (requestCode) {
            case 1: //REQUEST_ACCOUNT_PICKER
                if (data != null && data.getExtras() != null) {
                    String accountName = data.getExtras().getString(AccountManager.KEY_ACCOUNT_NAME);
                    if (accountName != null) {
                        user2.setSelectedAccountName(accountName);
                        userStatus.setText(user2.getSelectedAccountName());
                    }
                }
            break;
        }
    }

    static void toaster(Context c, String s) {
        Toast toast = Toast.makeText(c, s, Toast.LENGTH_SHORT);
        toast.setGravity(Gravity.TOP|Gravity.RIGHT, 0, 0);
        toast.show();
    }

}