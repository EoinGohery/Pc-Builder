package com.example.BriarfieldPC.demo.controller;

import com.example.BriarfieldPC.demo.entity.*;
import com.example.BriarfieldPC.demo.service.PcService;
import lombok.AllArgsConstructor;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/parts")
@AllArgsConstructor
public class PcController {
    @Autowired
    private PcService pcService;

    @GetMapping("/getGPUs")
    public List<Gpu> getAllGpus(){
       return pcService.getAllGpu();
    }
    @GetMapping("/getCPUs")
    public List<Cpu> getAllCpus(){
        return pcService.getAllCpu();
    }
    @GetMapping("/getDrivers")
    public List<Drive> getAllDriver(){
        return pcService.getAllDrivers();
    }
    @GetMapping("/getMBDs")
    public List<Motherboard> getAllMotherboard(){
        return pcService.getAllMotherboards();
    }
    @GetMapping("/getPS")
    public List<PSU> getAllPS(){
        return pcService.getAllPSU();
    }
    @GetMapping("/getRam")
    public List<Ram> getAllTheRam(){
        return pcService.getAllRam();
    }
}
