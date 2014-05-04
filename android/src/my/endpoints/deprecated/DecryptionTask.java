package my.endpoints.deprecated;
/*
import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.util.Pair;
import com.appspot.gcl_13.rest.Rest;
import com.google.api.client.extensions.android.http.AndroidHttp;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;
import com.google.api.client.googleapis.extensions.android.gms.auth.GooglePlayServicesAvailabilityIOException;
import com.google.api.client.googleapis.extensions.android.gms.auth.UserRecoverableAuthIOException;
import com.google.api.client.json.gson.GsonFactory;

import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.crypto.spec.SecretKeySpec;
import java.io.IOException;
import java.security.SecureRandom;

class DecryptionTask extends AsyncTask<Pair<Context,String>, Void, Pair<Context,String>> {

    @Override
    protected Pair doInBackground(Pair<Context,String>... arg) {
        String err="";
        Context c = arg[0].first;
        String k = arg[0].second;

        byte[] yourKey = new byte[0];
        //byte[] decodedData = decodeFile(yourKey, bytesOfYourFile);

        return Pair.create(c, err);
    }

    @Override
    protected void onPostExecute(Pair<Context,String> p) {
        MainActivity.toaster(p.first, p.second);
    }

    public static byte[] decodeFile(byte[] key, byte[] fileData) throws Exception
    {
        SecretKeySpec skeySpec = new SecretKeySpec(key, "AES");
        Cipher cipher = Cipher.getInstance("AES");
        cipher.init(Cipher.DECRYPT_MODE, skeySpec);

        byte[] decrypted = cipher.doFinal(fileData);

        return decrypted;
    }

}
*/
/*
        File file = new File(Environment.getExternalStorageDirectory() + File.separator + "your_folder_on_sd", "file_name");
        BufferedOutputStream bos = new BufferedOutputStream(new FileOutputStream(file));
        byte[] yourKey = generateKey("password");
        byte[] filesBytes = encodeFile(yourKey, yourByteArrayContainigDataToEncrypt);
        bos.write(fileBytes);
        bos.flush();
        bos.close();

    public static byte[] encodeFile(byte[] key, byte[] fileData) throws Exception
    {
        SecretKeySpec skeySpec = new SecretKeySpec(key, "AES");
        Cipher cipher = Cipher.getInstance("AES");
        cipher.init(Cipher.ENCRYPT_MODE, skeySpec);

        byte[] encrypted = cipher.doFinal(fileData);

        return encrypted;
    }

    public static byte[] generateKey(String password) throws Exception
    {
        byte[] keyStart = password.getBytes("UTF-8");

        KeyGenerator kgen = KeyGenerator.getInstance("AES");
        SecureRandom sr = SecureRandom.getInstance("SHA1PRNG", "Crypto");
        sr.setSeed(keyStart);
        kgen.init(128, sr);
        SecretKey skey = kgen.generateKey();
        return skey.getEncoded();
    }


 */
